package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc094/tasks/abc094_a

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if abs(a-b) <= d && abs(b-c) <= d || abs(a-c) <= d {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func abs(x int) int {
	if x <0 {
		x = -x
	}
	return x
}