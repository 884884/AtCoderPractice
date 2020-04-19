package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc123/tasks/abc123_a

func main() {
	// a := make([]int, 5)
	// for i := 0; i < 5; i++ {
	// 	fmt.Scan(&a[i])
	// }
	//
	// var k int
	// fmt.Scan(&k)
	//
	// for i := 0; i < 4; i++ {
	// 	for j := i + 1; j < 5; j++ {
	// 		if v := a[j] - a[i]; v > k {
	// 			fmt.Println(":(")
	// 			return
	// 		}
	// 	}
	// }
	// fmt.Println("Yay!")
	// 問題文の制約を活かせば、以下のように必要な計算を減らせる

	var a, b, c, d, e, k int
	fmt.Scan(&a, &b, &c, &d, &e, &k)

	if e-a > k {
		fmt.Println(":(")
	} else {
		fmt.Println("Yay!")
	}
}
