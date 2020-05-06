package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc154/tasks/abc154_b

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		fmt.Print("x")
	}
}
