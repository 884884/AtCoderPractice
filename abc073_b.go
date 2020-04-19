package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc073/tasks/abc073_b

func main() {
	var n int
	fmt.Scan(&n)

	var l, r, sum int
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r)
		sum += r - (l - 1)
	}

	fmt.Println(sum)
}
