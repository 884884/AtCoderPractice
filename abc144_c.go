package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc144/tasks/abc144_c

func main() {
	var n int
	fmt.Scan(&n)

	var a, b int
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		a = i
		b = n / i
	}

	fmt.Println(a + b - 2)
}
