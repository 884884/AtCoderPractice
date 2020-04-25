package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc077/tasks/abc077_b

func main() {
	var n int
	fmt.Scan(&n)

	var ans int
	for i := 1; i * i <= n; i++ {
		ans = i * i
	}
	fmt.Println(ans)
}