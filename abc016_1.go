package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc016/tasks/abc016_1

func main() {
	var m, d int
	fmt.Scan(&m, &d)

	if m%d == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
