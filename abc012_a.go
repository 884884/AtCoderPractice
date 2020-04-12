package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc012/tasks/abc012_1

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// 単純に入れ替えればいいだけ
	// c = a
	// a = b
	// b = c
	// fmt.Println(a,b)
	fmt.Println(b, a)
}
