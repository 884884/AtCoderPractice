package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc055/tasks/abc055_a

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(800*n - 200*(n/15))
}
