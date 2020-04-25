package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/arc055/tasks/arc055_a

func main() {
	var n int
	fmt.Scan(&n)

	ans := "1"
	for i := 0; i < n-1; i++ {
		ans += "0"
	}
	ans += "7"

	fmt.Println(ans)
}

