package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc089/tasks/abc089_b

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[string]bool, n)

	var x string
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		m[x] = true
	}

	if len(m) == 3{
		fmt.Println("Three")
	} else {
		fmt.Println("Four")
	}
}
