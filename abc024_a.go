package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc024/tasks/abc024_a

func main() {
	var a, b, c, k, s, t int
	fmt.Scan(&a, &b, &c, &k, &s, &t)

	ans := a*s + b*t
	if s+t >= k {
		ans -= c*(s+t)
	}

	fmt.Println(ans)
}
