package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc031/tasks/abc031_a
func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}

func main() {
	var a, d int
	fmt.Scan(&a, &d)

	fmt.Println(max((a+1)*d, a*(d+1)))
}
