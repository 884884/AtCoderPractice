package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc070/tasks/abc070_a

func main() {
	var N string
	fmt.Scan(&N)

	if N[0] == N[2]{
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
