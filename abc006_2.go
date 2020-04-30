package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc006/tasks/abc006_2

func main() {
	var n int
	fmt.Scan(&n)

	as := make([]int, n+3)
	as[1], as[2], as[3] = 0, 0, 1
	for i := 4; i <= n; i++ {
		as[i] = (as[i-1] + as[i-2] + as[i-3]) % 10007
	}

	fmt.Println(as[n])
}
