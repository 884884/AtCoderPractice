package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc141/tasks/abc141_a

func main() {
	var s string
	fmt.Scan(&s)

	wm := map[string]string{
		"Sunny": "Cloudy",
		"Cloudy": "Rainy",
		"Rainy": "Sunny",
	}

	fmt.Println(wm[s])
}
