package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc131/tasks/abc131_a

func main() {
	var s string
	fmt.Scan(&s)

	if s[0] == s[1] {
		fmt.Println("Bad")
		return
	}
	if s[1] == s[2] {
		fmt.Println("Bad")
		return
	}
	if s[2] == s[3] {
		fmt.Println("Bad")
		return
	}
	fmt.Println("Good")
}
