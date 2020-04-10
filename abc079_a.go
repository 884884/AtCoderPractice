package main

import "fmt"

// 以下の問題
// https://atcoder.jp/contests/abc079/tasks/abc079_a

func main() {
	var N string
	fmt.Scan(&N)

	// これは冗長すぎる
	// if N[0] == N[1] {
	// 	if N[1] == N[2] {
	// 		fmt.Println("Yes")
	// 	} else {
	// 		fmt.Println("No")
	// 	}
	// } else {
	// 	if N[1] == N[2] {
	// 		if N[2] == N[3] {
	// 			fmt.Println("Yes")
	// 		} else {
	// 			fmt.Println("No")
	// 		}
	// 	} else {
	// 		fmt.Println("No")
	// 	}
	// }

	if N[0] == N[1] && N[1] == N[2] {
		fmt.Println("Yes")
		return
	}
	if N[1] == N[2] && N[2] == N[3] {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}