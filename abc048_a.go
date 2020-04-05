package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc048/tasks/abc048_a

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)

	fmt.Println(string(s1[0]) + string(s2[0]) + string(s3[0]))
}
