package main

import (
	"fmt"
	"strconv"
)

// 以下の問題
// https://atcoder.jp/contests/abc086/tasks/abc086_b

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	num, _ := strconv.Atoi(a + b)

	for i := 1; i <= num; i++ {
		if i*i == num {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
