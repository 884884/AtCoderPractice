package main

import (
	"fmt"
	"sort"
	"strings"
)

// 以下の問題
// https://atcoder.jp/contests/abc082/tasks/abc082_b

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	// stringを並び替える方法
	// 1. strings.Splitで1文字ずつ[]stringにする
	// 2. []stringをsortする（sortするにはsliceではないとできないため）
	// 3. []stringをJoinしてstringにする
	s2 := strings.Split(s, "")
	t2 := strings.Split(t, "")

	sort.Strings(s2)
	sort.Sort(sort.Reverse(sort.StringSlice(t2)))

	s = strings.Join(s2, "")
	t = strings.Join(t2, "")

	if s < t {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
