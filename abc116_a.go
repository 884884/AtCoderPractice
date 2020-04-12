package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc116/tasks/abc116_a

func main() {
	var ab, bc int
	fmt.Scan(&ab, &bc)
	fmt.Println(bc * ab / 2)
}
