package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc118/tasks/abc118_b

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	count := make(map[int]int, M)
	var K, A int
	for i := 0; i < N; i++ {
		fmt.Scan(&K)
		for j := 0; j < K; j++ {
			fmt.Scan(&A)
			count[A] ++
		}
	}

	var ans int
	for _, v := range count {
		if v == N {
			ans++
		}
	}

	fmt.Println(ans)
}
