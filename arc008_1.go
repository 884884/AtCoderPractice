package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/arc008/tasks/arc008_1

func main() {
	var n int
	fmt.Scan(&n)

	ans := n / 10 * 100
	if n%10 < 7 {
		ans += n % 10 * 15
	} else {
		ans += 100
	}
	fmt.Println(ans)
}
