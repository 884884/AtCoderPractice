package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc066/tasks/abc066_b

func main() {
	var s string
	fmt.Scan(&s)

	for {
		s = s[:len(s)-1]
		if len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:] {
			fmt.Println(len(s))
			return
		}
	}
}
