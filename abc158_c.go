package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc158/tasks/abc158_c

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i <= 1009; i++ {
		if i*8/100 == a && i*1/10 == b {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
