package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc075/tasks/abc075_a

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b {
		fmt.Println(c)
	} else if b == c {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
