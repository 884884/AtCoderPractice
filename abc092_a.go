package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc092/tasks/abc092_a

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	// if a >= b{
	// 	train = b
	// } else {
	// 	train = a
	// }
	// if c >= d {
	// 	bus = d
	// } else {
	// 	bus = c
	// }
	// この書き方はイケてない。min関数を使う
	fmt.Println(math.Min(a,b)+math.Min(c,d))
}
