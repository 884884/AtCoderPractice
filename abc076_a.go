package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc076/tasks/abc076_a

func main() {
	var r, g int
	fmt.Scan(&r, &g)

	fmt.Println(2*g - r)
}
