package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc044/tasks/abc044_a

func main() {
	var n, k, x, y int
	fmt.Scan(&n, &k, &x, &y)
	if n > k {
		fmt.Println(x*n + y*(n-k))
	} else {
		fmt.Println(x * n)
	}
}
