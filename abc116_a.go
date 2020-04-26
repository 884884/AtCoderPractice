package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc116/tasks/abc116_a

func main() {
	var ab, bc, ca int
	fmt.Scan(&ab, &bc, &ca)

	fmt.Println(ab * bc / 2)
}
