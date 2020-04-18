package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc118/tasks/abc118_b

func main() {
	var n, m int
	fmt.Scan(&n)
	slice2d := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		slice2d[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&slice2d[i][j])
		}
	}
	fmt.Println(slice2d)
}