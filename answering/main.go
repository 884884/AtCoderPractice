package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc025/tasks/abc025_a

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)

	fmt.Println(s[n/5-1:n/5] + s[n%5-1:n%5])
}
