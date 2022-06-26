package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

const (
	// sampleEmoji
	// 絵文字が正しく表示される環境であれば1文字に見える
	// RuneCount すると 3つの絵文字とそれらをつなぐゼロ幅接合子を数える
	sampleEmoji = "🏊🏽‍♀️"
	sampleText  = "𠮷野家で𩸽たのんで𠮟られる🏊🏽‍♀️"
)

func main() {
	fmt.Println(sampleEmoji, "を数える")
	demonstrationCount()

	fmt.Println(sampleText, "を切り取る")
	fmt.Println(cut(14, sampleText))
}

func demonstrationCount() {
	fmt.Println("len():", countWithLen(sampleEmoji))
	fmt.Println("RuneCountInString():", countWithRuneCount(sampleEmoji))
	fmt.Println("GraphemeClusterCount():", countWithUniseg(sampleEmoji))
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

// uniseg が用意する graphemes を見た目上の文字ごとに取り出し
// 文字数に達するか、s が終わるまで数え、返却する
func cut(maxLength int, s string) string {
	var result []rune

	graphemes := uniseg.NewGraphemes(s)
	var i int = 0
	for graphemes.Next() && i < maxLength {
		result = append(result, graphemes.Runes()...)
		i++
	}

	return string(result)
}
