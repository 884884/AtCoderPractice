package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc096/tasks/abc096_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := a
	if b < a {
		ans--
	}

	fmt.Println(ans)
}
