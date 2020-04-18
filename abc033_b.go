package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc033/tasks/abc033_b

func main() {
	var N int
	fmt.Scan(&N)

	var (
		s        string
		p, total int
	)
	cities := make(map[string]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s, &p)
		cities[s] = p
		total += p
	}

	for i, v := range cities {
		if v >= total/2+1 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("atcoder")
}
