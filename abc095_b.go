package main

import (
	"fmt"
	"sort"
)

// 以下の問題
// https://atcoder.jp/contests/abc095/tasks/abc095_b

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	var sum int
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
		sum +=arr[i]
	}
	sort.Ints(arr)

	fmt.Println(len(arr) + (x-sum)/arr[0])
}
