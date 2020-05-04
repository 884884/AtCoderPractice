package main

import (
	"fmt"
	"strings"
)

// 以下の問題
// https://atcoder.jp/contests/abc073/tasks/abc073_a

func main() {
	var n string
	fmt.Scan(&n)

	if strings.Contains(n, "9") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
