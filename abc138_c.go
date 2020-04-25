package main

import (
	"fmt"
	"sort"
)

// 以下の問題
// https://atcoder.jp/contests/abc138/tasks/abc138_c

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]float64, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	sort.Float64s(arr)

	sum := arr[0]
	for i := 1; i < n; i++ {
		sum = (sum + arr[i]) / 2
	}

	fmt.Println(sum)
}
