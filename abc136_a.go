package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc136/tasks/abc136_a

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var ans int
	if ans = c - (a - b); ans < 0 {
		ans = 0
	}
	fmt.Println(ans)
}
