package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc121/tasks/abc121_b

func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)

	bArr := make([]int, m)
	for i := range bArr {
		fmt.Scan(&bArr[i])
	}

	var a, ans int
	for i := 0; i < n; i++ {
		total := c
		for j := 0; j < m; j++ {
			fmt.Scan(&a)
			total += a * bArr[j]
		}
		if total > 0 {
			ans++
		}
	}

	fmt.Println(ans)
}
