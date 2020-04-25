package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc055/tasks/abc055_b

func main() {
	var n int
	fmt.Scan(&n)

	power := 1
	for i := 1; i <= n; i++ {
		power *= i
		power %= 1000000007
	}

	fmt.Println(power)
}
