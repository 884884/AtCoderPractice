package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc138/tasks/abc138_a

func main() {
	var (
		a int
		s string
	)
	fmt.Scan(&a, &s)

	if a >= 3200 {
		fmt.Println(s)
	} else {
		fmt.Println("red")
	}
}
