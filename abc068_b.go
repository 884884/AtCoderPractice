package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc068/tasks/abc068_b

func main() {
	var n int
	fmt.Scan(&n)

	ans := 1
	for ans*2 <= n{
		ans *= 2
	}

	fmt.Println(ans)
}
