package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/tenka1-2019-beginner/tasks/tenka1_2019_a

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c > a && c < b || c < a && c > b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
