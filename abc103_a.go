package main

import (
	"fmt"
	"sort"
)

// 以下の問題
// https://atcoder.jp/contests/abc103/tasks/abc103_a

func main() {
	a := make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])
	sort.Ints(a)
	fmt.Println(a[2] - a[0])
}
