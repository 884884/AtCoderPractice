package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc122/tasks/abc122_b

func main() {
	var s string
	fmt.Scan(&s)

	var count, max int
	for _, v := range s {
		if v == 'A' || v == 'C' || v == 'G' || v == 'T' {
			count ++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
	}

	fmt.Println(max)
}
