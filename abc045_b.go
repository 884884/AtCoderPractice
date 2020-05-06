package main

import (
	"fmt"
)

// 以下の問題
// https://atcoder.jp/contests/abc045/tasks/abc045_b

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	x := a[0:1]
	a = a[1:]
	for {
		switch x {
		case "a":
			if a == "" {
				fmt.Println("A")
				return
			}
			x = a[0:1]
			a = a[1:]
		case "b":
			if b == "" {
				fmt.Println("B")
				return
			}
			x = b[0:1]
			b = b[1:]
		case "c":
			if c == "" {
				fmt.Println("C")
				return
			}
			x = c[0:1]
			c = c[1:]
		}
	}

}
