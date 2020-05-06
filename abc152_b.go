package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc152/tasks/abc152_b

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}

	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
}
