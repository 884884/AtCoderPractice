package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/arc002/tasks/arc002_1

func main() {
	var y int
	fmt.Scan(&y)

	if y%400 == 0 {
		fmt.Println("YES")
		return
	} else if y%100 == 0 {
		fmt.Println("NO")
		return
	} else if y%4 == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
