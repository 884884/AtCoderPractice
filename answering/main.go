package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc150/tasks/abc150_a

func main() {
	var k, x int
	fmt.Scan(&k, &x)

	if 500*k >= x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
