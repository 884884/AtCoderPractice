package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc087/tasks/abc087_a

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)

	fmt.Println((x - a) % b)
}
