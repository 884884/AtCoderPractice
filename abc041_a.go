package main

import "fmt"

func b() {
	var (
		s string
		i int
	)
	fmt.Scan(&s, &i)

	fmt.Println(s[i])
}
