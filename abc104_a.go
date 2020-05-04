package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc104/tasks/abc104_a

func main() {
	var r int
	fmt.Scan(&r)

	if r < 1200 {
		fmt.Println("ABC")
	} else if r < 2800 {
		fmt.Println("ARC")
	} else {
		fmt.Println("AGC")
	}
}
