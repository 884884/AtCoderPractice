package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc115/tasks/abc115_b

func main() {
	var N, p int
	fmt.Scan(&N)

	var max, add int
	for i := 0; i < N; i++ {
		fmt.Scan(&p)
		add += p
		if max < p {
			max = p
		}
	}

	fmt.Println(add - max/2)
}
