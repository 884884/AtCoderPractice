package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc065/tasks/abc065_a

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)

	if a >= b {
		fmt.Println("delicious")
	} else if b-a <= x {
		fmt.Println("safe")
	} else {
		fmt.Println("dangerous")
	}
}
