package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc146/tasks/abc146_b

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	var ans string
	for i := 0; i < len(s); i++ {
		a := s[i] + byte(n)
		if a > 'Z' {
			a = a - 26
		}
		ans += string(a)
	}

	fmt.Println(ans)
}
