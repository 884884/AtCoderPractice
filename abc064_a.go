package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc064/tasks/abc064_a

func main() {
	var r, g, b int
	fmt.Scan(&r, &g, &b)

	if (r*100+g*10+b)%4 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
