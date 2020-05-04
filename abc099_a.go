package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc099/tasks/abc099_a

func main() {
	var n int
	fmt.Scan(&n)

	if n <=999 {
		fmt.Println("ABC")
	} else {
		fmt.Println("ABD")
	}
}
