package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc138/tasks/abc138_b

func main() {
	var n int
	fmt.Scan(&n)

	var sum, a float64
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum += 1 / a
	}

	fmt.Println(1 / sum)
}
