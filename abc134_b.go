package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc134/tasks/abc134_b

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	fmt.Println(math.Ceil(float64(n) / float64(d*2+1)))
}
