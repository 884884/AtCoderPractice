package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc097/tasks/abc097_b

func main() {
	var x int
	fmt.Scan(&x)

	ans := 1
	for b := 2; b < x; b++ {
		for q := 2; q < x; q++ {
			num := int(math.Pow(float64(b), float64(q)))
			if num > x {
				break
			} else if num > ans {
				ans = num
			}
		}
	}

	fmt.Println(ans)
}
