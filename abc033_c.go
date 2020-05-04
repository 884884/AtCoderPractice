package main

import (
	"fmt"
	"strings"
)

// 以下の問題
// https://atcoder.jp/contests/abc033/tasks/abc033_c

func main() {
	var s string
	fmt.Scan(&s)

	ss := strings.Split(s, "+")

	var c int
	for _, v := range ss {
		if !strings.Contains(v, "0") {
			c++
		}
	}

	fmt.Println(c)
}
