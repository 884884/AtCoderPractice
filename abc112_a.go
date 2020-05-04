package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc112/tasks/abc112_a

func main() {
	var n int
	fmt.Scan(&n)

	var a, b int
	if n == 1 {
		fmt.Println("Hello World")
	} else {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
