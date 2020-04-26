package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc015/tasks/abc015_2

func main() {
	var n int
	fmt.Scan(&n)

	var sum, count, a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a != 0 {
			sum += a
			count ++
		}
	}

	fmt.Println(math.Ceil(float64(sum) / float64(count)))
}
