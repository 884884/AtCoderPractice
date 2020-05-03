package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc059/tasks/abc059_b

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	if a == b {
		fmt.Println("EQUAL")
	} else if len(a) >= len(b) && a > b {
		fmt.Println("GREATER")
	} else {
		fmt.Println("LESS")
	}
}
