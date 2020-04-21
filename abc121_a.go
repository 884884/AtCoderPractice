package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc121/tasks/abc121_a

func main() {
	var H, W, h, w int
	fmt.Scan(&H, &W, &h, &w)

	fmt.Println((H - h) * (W - w))
}
