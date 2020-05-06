package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc119/tasks/abc119_b

func main() {
	var n int
	fmt.Scan(&n)

	var (
		x, ans float64
		u      string
	)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &u)
		if u == "BTC" {
			x = 380000 * x
		}
		ans += x
	}

	fmt.Println(ans)
}
