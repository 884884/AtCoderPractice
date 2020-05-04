package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc136/tasks/abc136_c

func main() {
	var n int
	fmt.Scan(&n)

	hs := make([]int, n)
	for i := range hs {
		fmt.Scan(&hs[i])
	}

	for i := len(hs) - 2; i >= 0; i-- {
		switch {
		case hs[i]-hs[i+1] > 1:
			fmt.Println("No")
			return
		case hs[i]-hs[i+1] == 1:
			hs[i] -= 1
		}
	}
	fmt.Println("Yes")
}
