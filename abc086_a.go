package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc086/tasks/abc086_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if (a*b)%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
