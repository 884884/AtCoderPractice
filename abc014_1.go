package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc014/tasks/abc014_1

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a%b == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(b - (a % b))
	}
}
