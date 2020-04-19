package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc074/tasks/abc074_a

func main() {
	var n, a int
	fmt.Scan(&n, &a)

	fmt.Println(n*n-a)
}
