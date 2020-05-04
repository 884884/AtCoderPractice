package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc083/tasks/abc083_a

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if ab, cd := a+b, c+d; ab == cd {
		fmt.Println("Balanced")
	}else if ab > cd {
		fmt.Println("Left")
	}else {
		fmt.Println("Right")
	}
}
