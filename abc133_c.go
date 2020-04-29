package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc133/tasks/abc133_c

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	var min int
	ans := 2020

	for i := l; i < r; i++ {
		for j := i + 1; j <= r; j++ {
			min = i * j % 2019
			if min < ans {
				ans = min
			}
			if min == 0 {
				fmt.Println(0)
				return
			}
		}
	}

	fmt.Println(ans)
}
