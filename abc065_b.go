package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc065/tasks/abc065_b

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n+1)
	arr[0] = -1
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	c, v := 0, 1
	for i := 1; i <= n; i++ {
		if v == 2 {
			fmt.Println(c)
			return
		}
		v = arr[v]
		c++
	}

	fmt.Println("-1")
}
