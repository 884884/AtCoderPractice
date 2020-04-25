package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc039/tasks/abc039_b

func main() {
	var x int
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		if i*i*i*i == x {
			fmt.Println(i)
			return
		}
	}

}
