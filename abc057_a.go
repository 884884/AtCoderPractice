package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc057/tasks/abc057_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println((a + b) % 24)
}
