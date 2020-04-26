package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc108/tasks/abc108_a

func main() {
	var k int
	fmt.Scan(&k)

	odd := k/2 + k%2
	even := k/2
	fmt.Println(odd * even)
}
