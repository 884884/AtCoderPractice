package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc036/tasks/abc036_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println((b + a - 1) / a)
}
