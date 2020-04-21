package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc028/tasks/abc028_a

func main() {
	var n int
	fmt.Scan(&n)

	var ans string
	switch {
	case n <= 59:
		ans = "Bad"
	case 60 <= n && n <= 89:
		ans = "Good"
	case 90 <= n && n <= 99:
		ans = "Great"
	case n == 100:
		ans = "Perfect"
	}

	fmt.Println(ans)
}
