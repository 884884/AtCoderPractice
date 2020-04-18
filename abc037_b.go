package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc037/tasks/abc037_b

func main() {
	var N, Q int
	fmt.Scan(&N, &Q)

	sequence := make([]int, N)

	var L, R, T int
	for i := 0; i < Q; i++ {
		fmt.Scan(&L, &R, &T)
		for i := L - 1; i <= R-1; i++ {
			sequence[i] = T
		}
	}

	for _, v := range sequence {
		fmt.Println(v)
	}
}
