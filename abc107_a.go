package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc107/tasks/abc107_a

func main() {
	var n, i int
	fmt.Scan(&n, &i)

	fmt.Println(n - i + 1)
}