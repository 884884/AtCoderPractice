package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc067/tasks/abc067_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a%3 == 0 || b%3 == 0 || (a+b)%3 == 0 {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}
}
