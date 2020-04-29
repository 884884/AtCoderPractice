package main

import (
	"fmt"
	"strings"
)

// 以下の問題
// https://atcoder.jp/contests/abc081/tasks/abc081_a

func main() {
	var s string
	fmt.Scan(&s)

	// var ans int
	// for _ , v := range s {
	// 	if v == '1'{
	// 		ans ++
	// 	}
	// }

	// Countを使うのが一番簡単
	fmt.Println(strings.Count(s,"1"))
}
