package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

const (
	// sampleText
	// 絵文字が正しく表示される環境であれば1文字に見える
	// RuneCount すると 3つの絵文字とそれらをつなぐゼロ幅接合子を数える
	sampleText = "🏊🏽‍♀️"
)

func main() {
	fmt.Println(sampleText)
	demonstrationCount()

}

func demonstrationCount() {
	fmt.Println("len():", countWithLen(sampleText))
	fmt.Println("RuneCountInString():", countWithRuneCount(sampleText))
	fmt.Println("GraphemeClusterCount():", countWithUniseg(sampleText))
}

func countWithLen(s string) int {
	return len(s)
}

func countWithRuneCount(s string) int {
	return utf8.RuneCountInString(s)
}

func countWithUniseg(s string) int {
	return uniseg.GraphemeClusterCount(s)
}
