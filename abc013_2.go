package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc013/tasks/abc013_2

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	fmt.Println(math.Min(math.Min(math.Abs(a-b), math.Abs(10+b-a)), math.Abs(10+a-b)))
}
