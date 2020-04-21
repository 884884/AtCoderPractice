package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc052/tasks/abc052_a
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(max(a*b, c*d))
}
