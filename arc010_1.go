package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/arc010/tasks/arc010_1

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m, &a, &b)

	var c int
	for i := 1; i <= m; i++ {
		fmt.Scan(&c)
		if n <= a {
			n += b
		}
		if n - c < 0 {
			fmt.Println(i)
			return
		} else {
			n -= c
		}
	}
	fmt.Println("complete")
}
