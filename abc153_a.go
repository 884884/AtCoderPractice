package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc153/tasks/abc153_a
// https://daeudaeu.com/programming/c-language/kurisuke_kiriage_shisyagonyu/

func main() {
	var h, a int
	fmt.Scan(&h, &a)

	fmt.Println((h + a - 1)/a)
}
