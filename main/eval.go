package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 计算准确率、召回率和 F1 值
func evaluate(referencePath, candidatePath string) (float64, float64, float64) {
	refFile, err := os.Open(referencePath)
	if err != nil {
		fmt.Println("Error opening reference file:", err)
		return 0, 0, 0
	}
	defer refFile.Close()

	canFile, err := os.Open(candidatePath)
	if err != nil {
		fmt.Println("Error opening candidate file:", err)
		return 0, 0, 0
	}
	defer canFile.Close()

	refScanner := bufio.NewScanner(refFile)
	canScanner := bufio.NewScanner(canFile)

	var refCount, canCount, accCount int

	for refScanner.Scan() && canScanner.Scan() {
		refLine := refScanner.Text()
		canLine := canScanner.Text()

		refWords := strings.Split(refLine, " ")
		canWords := strings.Split(canLine, " ")

		refCount += len(refWords)
		canCount += len(canWords)

		// 计算正确分词的词数
		for _, refWord := range refWords {
			for _, canWord := range canWords {
				if refWord == canWord {
					accCount++
					break
				}
			}
		}
	}

	// 计算准确率、召回率和 F1 值
	P := float64(accCount) / float64(canCount) * 100
	R := float64(accCount) / float64(refCount) * 100
	F1 := (2 * P * R) / (P + R)

	return P, R, F1
}
