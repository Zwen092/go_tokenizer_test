package main

import (
	"github.com/huichen/sego"
	"github.com/wangbin/jiebago"
	"github.com/yanyiwu/gojieba"
	"reflect"
	"strings"
	"testing"

	"github.com/go-ego/gse"
)

// 测试 gojieba 的 Cut 方法
func TestGoJiebaCut(t *testing.T) {
	jieba := gojieba.NewJieba()
	defer jieba.Free()

	text := "人们常说生活是一部教科书"
	expected := []string{"人们", "常说", "生活", "是", "一部", "教科书"}

	result := jieba.Cut(text, true)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("gojieba Cut() failed, expected %v, got %v", expected, result)
	}
}

// 测试 jiebago 的 Cut 方法
func TestJiebaGoCut(t *testing.T) {
	jieba := &jiebago.Segmenter{}
	err := jieba.LoadDictionary("../jieba_dict.txt")
	if err != nil {
		t.Fatalf("Failed to load jiebago dictionary: %v", err)
	}

	text := "人们常说生活是一部教科书"
	expected := []string{"人们", "常说", "生活", "是", "一部", "教科书"}

	var result []string
	segments := jieba.Cut(text, true)
	for segment := range segments {
		result = append(result, segment)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("jiebago Cut() failed, expected %v, got %v", expected, result)
	}
}

// failed
func TestSegoCut(t *testing.T) {
	seg := sego.Segmenter{}
	seg.LoadDictionary("../sego_dict.txt")

	text := "人们常说生活是一部教科书"
	expected := []string{"人们", "常说", "生活", "是", "一部", "教科书"}

	segments := seg.Segment([]byte(text))
	result := sego.SegmentsToString(segments, false)
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

	if !reflect.DeepEqual(cleanedWords, expected) {
		t.Errorf("sego Cut() failed, expected %v, got %v", expected, cleanedWords)
	}
}

// 测试 gse 的 Cut 方法
func TestGseCut(t *testing.T) {
	seg := gse.Segmenter{}
	err := seg.LoadDict()
	if err != nil {
		return
	}

	text := "人们常说生活是一部教科书"
	expected := []string{"人们", "常说", "生活", "是", "一部", "教科书"}

	result := seg.Cut(text, true)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("gse Cut() failed, expected %v, got %v", expected, result)
	}
}
