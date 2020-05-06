package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc156/tasks/abc156_b

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	ans := 0
	for n != 0 {
		n /= k
		ans++
	}
	fmt.Println(ans)
}
