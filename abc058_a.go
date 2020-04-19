package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc058/tasks/abc058_a

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if b-a == c-b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
