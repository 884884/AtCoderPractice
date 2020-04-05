package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc041/submissions/me

func main() {
	var (
		s string
		i int
	)
	fmt.Scan(&s, &i)

	// 1バイト（シングルバイト）文字なら以下でOK
	// fmt.Println(s[i-1],s[i])

	// 1バイトより多いマルチバイト文字なら,rune型で扱う
	// []rune()でstringからruneへキャストする
	runeSlice := []rune(s)
	fmt.Println(string(runeSlice[i-1]))
}
