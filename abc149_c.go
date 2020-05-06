package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc149/tasks/abc149_c

func main() {
	var x int
	fmt.Scan(&x)

	for i := x; i <= 100003; i++ {
		ok := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(i)
			return
		}
	}
}
