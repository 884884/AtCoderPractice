package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc109/tasks/abc109_a

func main() {
	var x int
	fmt.Scan(&x )

	if x == 3 || x == 5 || x == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
