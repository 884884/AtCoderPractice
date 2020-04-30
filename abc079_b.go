package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc079/tasks/abc079_b

func main() {
	var n int
	fmt.Scan(&n)

	as := make([]int, n+1)
	as[0], as[1] = 2, 1
	for i := 2; i <= n; i++ {
		as[i] = as[i-1] + as[i-2]
	}

	fmt.Println(as[n])
}
