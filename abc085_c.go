package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc085/tasks/abc085_c

func main() {
	var N, Y int
	fmt.Scan(&N, &Y)

	for i := 0; i <= N; i++ {
		for j := 0; j <= N-i; j++ {
			k := N - j - i
			if 10000*i+5000*j+1000*k == Y {
				fmt.Println(i, j, k)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
