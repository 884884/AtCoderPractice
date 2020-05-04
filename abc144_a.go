package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc144/tasks/abc144_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a <= 9 && b <= 9 {
		fmt.Println(a * b)
	} else {
		fmt.Println(-1)
	}
}
