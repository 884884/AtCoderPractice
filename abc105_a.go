package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc105/tasks/abc105_a

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if n%k == 0{
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
