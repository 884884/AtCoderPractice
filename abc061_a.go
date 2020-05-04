package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc061/tasks/abc061_a

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c >= a && c <= b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
