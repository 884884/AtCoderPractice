package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc100/tasks/abc100_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a <= 8 && b <= 8 {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
