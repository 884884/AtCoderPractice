package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc142/tasks/abc142_a

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(float64(n - n/2) / float64(n))
}
