package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc106/tasks/abc106_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(a*b - (a + b - 1))
}