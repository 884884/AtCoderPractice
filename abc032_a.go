package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc032/tasks/abc032_a

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	for i := n;; i++ {
		if i%a == 0 && i%b == 0 {
			fmt.Println(i)
			return
		}
	}
}
