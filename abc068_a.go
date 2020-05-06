package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc068/tasks/abc068_a

func main() {
	var c1, c2, c3 string
	fmt.Scan(&c1, &c2, &c3)

	fmt.Println(c1[0:1] + c2[1:2] + c3[2:3])
}
