package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/arc018/tasks/arc018_1

func main() {
	var h, b float64
	fmt.Scan(&h, &b)

	fmt.Println(h/100 * h/100 *b)
}
