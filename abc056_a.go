package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc056/tasks/abc056_a

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	// if a == "H" {
	// 	if b == "H" {
	// 		fmt.Println("H")
	// 	} else {
	// 		fmt.Println("D")
	// 	}
	// } else {
	// 	if b == "H" {
	// 		fmt.Println("D")
	// 	} else {
	// 		fmt.Println("H")
	// 	}
	// }
	// これは冗長すぎる

	if a == b {
		fmt.Println("H")
	} else {
		fmt.Println("D")
	}
}
