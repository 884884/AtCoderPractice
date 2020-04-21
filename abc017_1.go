package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc017/tasks/abc017_1

func main() {
	var s, e, sum int
	for i := 0; i < 3; i++ {
		fmt.Scan(&s, &e)
		sum += s * e / 10
	}

	fmt.Println(sum)
}
