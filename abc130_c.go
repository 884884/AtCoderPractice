package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc130/tasks/abc130_c

func main() {
	var w, h, x, y float64
	fmt.Scan(&w, &h, &x, &y)

	ans := 0
	if x == w/2 && y == h/2 {
		ans = 1
	}
	fmt.Println(w*h/2, ans)
}
