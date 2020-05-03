package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc119/tasks/abc119_a

func main() {
	var s string
	fmt.Scan(&s)

	if s <= "2019/04/30" {
		fmt.Println("Heisei")
	} else {
		fmt.Println("TBD")
	}
}
