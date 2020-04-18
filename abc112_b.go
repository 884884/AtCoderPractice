package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc112/tasks/abc112_b

func main() {
	var N, T int
	fmt.Scan(&N, &T)

	var c, t int
	min := 9999
	for i := 0; i < N; i++ {
		fmt.Scan(&c, &t)
		if t <= T && c <= min {
			min = c
		}
	}

	if min != 9999 {
		fmt.Println(min)
		return
	}
	fmt.Println("TLE")
}
