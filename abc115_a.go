package main

import (
	"fmt"
	"strings"
)

// 以下の問題
// https://atcoder.jp/contests/abc115/tasks/abc115_a

func main() {
	var d int
	fmt.Scan(&d)

	fmt.Println("Christmas",strings.Repeat("Eve ",25-d))
}
