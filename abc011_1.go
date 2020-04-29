package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc011/tasks/abc011_1

func main() {
	var n int
	fmt.Scan(&n)

	if n == 12 {
		fmt.Println(1)
	} else {
		fmt.Println(n + 1)
	}
}
