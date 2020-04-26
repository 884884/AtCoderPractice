package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc123/tasks/abc123_b

func main() {
	var f, last, ans int

	fmt.Scan(&f)
	last = f
	for i := 0; i < 4; i++ {
		fmt.Scan(&f)
		if f%10 == 0 {
			ans += f
		} else if last%10 > f%10 {
			ans += ((last + 9) / 10) * 10
			last = f
		} else {
			ans += ((f + 9) / 10) * 10
		}
	}
	ans += last

	fmt.Println(ans)
}
