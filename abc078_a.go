package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc078/tasks/abc078_a

func main() {
	var x, y string
	fmt.Scan(&x, &y)

	if x == y {
		fmt.Println("=")
	} else if x < y {
		fmt.Println("<")
	} else {
		fmt.Println(">")
	}
}
