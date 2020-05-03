package main

import (
"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc130/tasks/abc130_a

func main() {
	var x, a int
	fmt.Scan(&x, &a)

	ans := 10
	if x < a {
		ans = 0
	}

	fmt.Println(ans)
}