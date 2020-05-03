package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc016/tasks/abc016_2

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	ans := "!"
	switch {
	case a+b == c && a-b == c:
		ans = "?"
	case a+b == c:
		ans = "+"
	case a-b == c:
		ans = "-"
	}

	fmt.Println(ans)
}
