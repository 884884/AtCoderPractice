package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc102/tasks/abc102_a

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n)
	} else {
		fmt.Println(n * 2)
	}
}
