package main

import (
	"fmt"
	"math"
)

// 以下の問題
// https://atcoder.jp/contests/abc041/tasks/abc041_b

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	divisor := int(math.Pow(10, 9) + 7)
	fmt.Println(a * b % divisor * c % divisor)
}
