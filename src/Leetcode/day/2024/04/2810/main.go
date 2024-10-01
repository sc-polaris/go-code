package main

import (
	"slices"
)

func finalString(s string) string {
	q := [2][]rune{} // 两个 slice 背靠背，q[0] 向左，q[1] 向右
	dir := 1
	for _, c := range s {
		if c == 'i' {
			dir ^= 1 // 添加修改方向
		} else {
			q[dir] = append(q[dir], c)
		}
	}
	slices.Reverse(q[dir^1])
	return string(append(q[dir^1], q[dir]...))
}
