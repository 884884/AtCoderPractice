package main

import (
	"fmt"
	"strings"
)

// 以下の問題
// https://atcoder.jp/contests/abc051/tasks/abc051_a

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(strings.Replace(s,","," ", -1))
}
