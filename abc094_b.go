package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc094/tasks/abc094_b

func main() {
	var N, M, X int
	fmt.Scan(&N, &M, &X)

	var first, last float64
	for i := 0; i < M; i++ {
		var A int
		fmt.Scan(&A)
		if A > X {
			last ++
		} else {
			first ++
		}
	}
	fmt.Println(math.Min(last,first))
}
