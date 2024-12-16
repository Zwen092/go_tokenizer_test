package main

import (
	"bufio"
	"fmt"
	"github.com/go-ego/gse"
	"github.com/huichen/sego"
	"github.com/wangbin/jiebago"
	"github.com/yanyiwu/gojieba"
	"os"
	"strings"
	"time"
)

// 分词函数类型
type SegmentFunc func(string) []string

func segmentAndMesureTime(segmentFunc SegmentFunc, text string) ([]string, time.Duration) {
	start := time.Now() // 记录开始时间
	result := segmentFunc(text)
	end := time.Now()             // 记录结束时间
	return result, end.Sub(start) // 返回分词结果和耗时
}

// 使用 gojieba 分词
func segmentGoJieba(jieba *gojieba.Jieba) SegmentFunc {
	return func(text string) []string {
		return jieba.Cut(text, true)
	}
}

// 使用 jiebago 分词
func segmentJiebaGo(jieba *jiebago.Segmenter) SegmentFunc {
	return func(text string) []string {
		var words []string
		segments := jieba.Cut(text, true)
		for segment := range segments {
			words = append(words, segment)
		}
		return words
	}
}

// 使用 sego 分词
func segmentSego(seg sego.Segmenter) SegmentFunc {
	return func(text string) []string {
		segments := seg.Segment([]byte(text))
		// 将分词结果转换为字符串

		result := sego.SegmentsToString(segments, false)
		// 将结果按 `/` 拆分为字符串数组
		words := strings.Split(result, " ")
		// 进一步处理，去除 `/` 和词性标注
		var cleanedWords []string
		for _, word := range words {
			// 按 `/` 拆分
			parts := strings.Split(word, "/")
			if len(parts) > 0 {
				cleanedWords = append(cleanedWords, parts[0]) // 只取词，忽略词性标注
			}
		}
		return cleanedWords
	}
}

// 使用 gse 分词
func segmentGse(seg gse.Segmenter) SegmentFunc {
	return func(text string) []string {
		segments := seg.Cut(text, true)
		return segments
	}
}

// 分词测试函数
func segmentTest(segmentFunc SegmentFunc, testFilePath, outputFilePath string) time.Duration {
	// 打开输入文件
	testFile, err := os.Open(testFilePath)
	if err != nil {
		fmt.Println("Error opening test file:", err)
		return 0
	}
	defer testFile.Close()

	// 打开输出文件
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return 0
	}
	defer outputFile.Close()

	// 逐行读取测试文件并分词
	scanner := bufio.NewScanner(testFile)
	var totalTime time.Duration
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			outputFile.WriteString("\n")
			continue
		}
		words, duration := segmentAndMesureTime(segmentFunc, line)
		totalTime += duration
		outputFile.WriteString(strings.Join(words, " ") + "\n")
	}
	return totalTime
}

func main() {

	// 输入文件路径
	testFilePath := "data/msr_test.txt"
	referenceFilePath := "data/msr.txt"

	// 初始化分词器
	jiebaGo := &jiebago.Segmenter{}
	err := jiebaGo.LoadDictionary("jieba_dict.txt")
	if err != nil {
		fmt.Println("Error loading jiebago dictionary:", err)
		return
	}

	segoSeg := sego.Segmenter{}
	segoSeg.LoadDictionary("sego_dict.txt")

	gseSeg := gse.Segmenter{}
	err = gseSeg.LoadDict()
	if err != nil {
		return
	}

	gojiebaSeg := gojieba.NewJieba()
	defer gojiebaSeg.Free()

	// 分词器列表
	segmenters := map[string]SegmentFunc{
		"gojieba": segmentGoJieba(gojiebaSeg),
		"gse":     segmentGse(gseSeg),
		"jiebago": segmentJiebaGo(jiebaGo),
		"sego":    segmentSego(segoSeg),
	}

	// 测试每个分词器
	for name, segmentFunc := range segmenters {
		fmt.Printf("Testing %s...\n", name)

		// 输出文件路径
		outputFilePath := fmt.Sprintf("report/msr_result_%s.txt", name)

		// 分词测试
		totalTime := segmentTest(segmentFunc, testFilePath, outputFilePath)

		// 评估分词结果
		P, R, F1 := evaluate(referenceFilePath, outputFilePath)
		fmt.Printf("Results for %s:\n", name)
		fmt.Printf("准确率: %.2f%%\n", P)
		fmt.Printf("召回率: %.2f%%\n", R)
		fmt.Printf("F1: %.2f%%\n", F1)
		fmt.Printf("总耗时: %v\n", totalTime)
		fmt.Println()
	}
}
