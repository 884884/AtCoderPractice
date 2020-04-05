package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/tenka1-2019-beginner/tasks/tenka1_2019_b

func main() {
	var (
		n, k int
		s, t string
	)
	fmt.Scan(&n, &s, &k)
	for i := 0; i < n; i++ {
		if s[i] != s[k-1] {
			t += "*"
		} else {
			t += string(s[i])
		}
	}
	fmt.Println(t)
}
