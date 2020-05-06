package main

import (
	"fmt"
	"strconv"
)

// 以下の問題
// https://atcoder.jp/contests/abc020/tasks/abc020_b

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	ans, _ := strconv.Atoi(a + b)
	fmt.Println(ans * 2)
}
