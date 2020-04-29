package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc158/tasks/abc158_b

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	forward := n / (a + b) * a
	back := n % (a + b)
	if a < back{
		back = a
	}

	fmt.Println(forward + back)
}
