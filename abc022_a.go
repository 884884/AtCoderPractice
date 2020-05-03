package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc022/tasks/abc022_a

func main() {
	var n, s, t, a int
	fmt.Scan(&n, &s, &t)
	c, w := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		w += a
		if w >= s && w <= t {
			c++
		}
	}

	fmt.Println(c)
}
