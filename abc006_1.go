package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc006/tasks/abc006_1

func main() {
	var n int
	fmt.Scan(&n)

	ans := "NO"
	if n % 3 == 0 {
		ans = "YES"
	}

	fmt.Println(ans)
}
