package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc081/tasks/abc081_b

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	ans := 0
	for {
		for i, v := range arr {
			if v%2 == 0 {
				arr[i] = v / 2
			} else {
				fmt.Println(ans)
				return
			}
		}
		ans++
	}
}
