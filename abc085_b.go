package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc085/tasks/abc085_b

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}

	fmt.Println(len(removeDuplicate(arr)))
}

func removeDuplicate(args []int) []int {
	results := make([]int, 0, len(args))
	check := map[int]bool{}
	for _, v := range args {
		// faleであれば、if文内を実行する
		if !check[v] {
			check[v] = true
			results = append(results, v)
		}
	}
	return results
}

// 重複チェックしないやり方だとこういう書き方でも行ける
// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	m := make(map[int]int)
// 	for i := 0; i < n; i++ {
// 		var s int
// 		fmt.Scan(&s)
// 		m[s]++
// 	}
//
// 	fmt.Println(len(m))
// }