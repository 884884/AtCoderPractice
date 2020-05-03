package main

import (
	"fmt"
	"strings"
)

// 以下の問題
// https://atcoder.jp/contests/abc029/tasks/abc029_b

func main() {
	var s string
	count := 0
	for i:= 0 ; i < 12 ; i++{
		fmt.Scan(&s)
		if strings.Contains(s,"r"){
			count ++
		}
	}

	fmt.Println(count)
}
