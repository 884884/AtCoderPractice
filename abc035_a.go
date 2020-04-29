package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc035/tasks/abc035_a

func main() {
	var w, h int
	fmt.Scan(&w, &h)

	if w%16 == 0 && h%9 == 0 {
		fmt.Println("16:9")
	} else {
		fmt.Println("4:3")
	}
}
