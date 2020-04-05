package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc058/tasks/abc058_b

func main() {
	var (
		o, e, a string
		diff bool
		)
	fmt.Scan(&o, &e)

	if len(o) != len(e){
		diff = true
	}

	for i := 0; i < len(o); i++ {
		a += string(o[i])
		if i == len(o)-1 && diff == true{
			break
		}
		a += string(e[i])
	}
	fmt.Println(a)
}
