package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc031/tasks/abc031_b

func main() {
	var L, H, N int
	fmt.Scan(&L, &H, &N)

	times := make([]int, N)
	for i := range times {
		fmt.Scan(&times[i])
	}

	var ans int
	for _, v := range times {
		switch {
		case v < L:
			ans = L - v
		case v >= L && v <= H:
			ans = 0
		case v > H:
			ans = -1
		}
		fmt.Println(ans)
	}
}
