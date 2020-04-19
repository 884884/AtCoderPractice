package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc072/tasks/abc072_a

func main() {
	var x, t int
	fmt.Scan(&x, &t)

	if v := x - t; v > 0 {
		fmt.Println(v)
	} else {
		fmt.Println(0)
	}
}
