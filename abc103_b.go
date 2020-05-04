package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc103/tasks/abc103_b

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	for i := 0; i < len(s); i++ {
		if s == t {
			fmt.Println("Yes")
			return
		}
		s = s[len(s)-1:] + s[:len(s)-1]
	}
	fmt.Println("No")
}
