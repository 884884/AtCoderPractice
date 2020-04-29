package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc034/tasks/abc034_b

func main() {
	var n int
	fmt.Scan(&n)

	var ans int
	if n % 2 == 0 {
		ans = n - 1
	} else {
		ans = n + 1
	}

	fmt.Println(ans)
}
