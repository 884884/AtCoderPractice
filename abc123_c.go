package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc123/tasks/abc123_c

func main() {
	var n, x int
	fmt.Scan(&n)

	fmt.Scan(&x)
	min := x
	for i := 0; i < 4; i++ {
		fmt.Scan(&x)
		if x < min {
			min = x
		}
	}

	fmt.Println(int(math.Ceil(float64(n)/float64(min))) + 4)
}

ABC153 A - Serval vs Monster
小数点切り上げ除算