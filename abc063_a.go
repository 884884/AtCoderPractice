package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc063/tasks/abc063_a

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if v := a+b ; v < 10{
		fmt.Println(v)
	} else {
		fmt.Println("error")
	}
}
