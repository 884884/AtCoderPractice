package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc030/tasks/abc030_a

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var ans string
	switch {
	case b*c > d*a:
		ans = "TAKAHASHI"
	case b*c < d*a:
		ans = "AOKI"
	default:
		ans = "DRAW"
	}
	fmt.Println(ans)
}
