package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc053/tasks/abc053_a

func main() {
	var x int
	fmt.Scan(&x)

	if x < 1200 {
		fmt.Println("ABC")
	} else {
		fmt.Println("ARC")
	}
}
