package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc034/tasks/abc034_a

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x < y {
		fmt.Println("Better")
	} else {
		fmt.Println("Worse")
	}
}
