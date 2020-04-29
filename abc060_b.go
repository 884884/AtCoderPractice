package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc060/tasks/abc060_b

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	for i := 0; i < b; i++ {
		if a*i%b == c {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
