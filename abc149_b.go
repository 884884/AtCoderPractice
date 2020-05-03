package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc149/tasks/abc149_b

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	ansT, ansA := a, b

	if a >= k {
		ansT -= k
	} else if b >=k-a {
		ansT = 0
		ansA -= k-a
	} else {
		ansT = 0
		ansA = 0
	}
	fmt.Println(ansT,ansA)
}
