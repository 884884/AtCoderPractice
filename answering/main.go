package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc144/tasks/abc144_c

func main() {
	var n int
	fmt.Scan(&n)

	ans := n - 1
	for i := 1; i <= n/2+1; i++ {
		if n%i != 0 {
			continue
		}

	}

	fmt.Println()
}
