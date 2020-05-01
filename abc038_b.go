package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc038/tasks/abc038_b

func main() {
	var h1, w1, h2, w2 int
	fmt.Scan(&h1, &w1, &h2, &w2)

	if h1 == h2 || h1 == w2 || w1 == h2 || w1 == w2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
