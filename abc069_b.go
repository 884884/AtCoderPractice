package main

import (
	"fmt"
	"strconv"
)

// 以下の問題
// https://atcoder.jp/contests/abc069/tasks/abc069_b

func a() {
	var s string
	fmt.Scan(&s)

	fmt.Println(s[:1] + strconv.Itoa(len(s)-2) + s[len(s)-1:])
}
