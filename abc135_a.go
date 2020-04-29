package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc135/tasks/abc135_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if (a+b)%2 == 0 {
		fmt.Println((a + b) / 2)
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}
