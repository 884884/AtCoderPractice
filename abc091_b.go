package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc091/tasks/abc091_b

func main() {
	var (
		N, M int
		s, t string
	)
	var st = map[string]int{}

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s)
		st[s]++
	}

	fmt.Scan(&M)
	for i := 0; i < M; i++ {
		fmt.Scan(&t)
		st[t]--
	}

	max := 0
	for _, v := range st {
		if max < v {
			max = v
		}
	}

	fmt.Println(max)
}
