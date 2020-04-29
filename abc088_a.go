package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc088/tasks/abc088_a

func main() {
	var n, a int
	fmt.Scan(&n, &a)

	if n%500 <= a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
