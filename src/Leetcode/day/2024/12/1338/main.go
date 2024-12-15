package main

import (
	"slices"
)

/*
	给你一个整数数组 arr。你可以从中选出一个整数集合，并删除这些整数在数组中的每次出现。

	返回 至少 能删除数组中的一半整数的整数集合的最小大小。
*/

func minSetSize(arr []int) int {
	cnt := make([]int, slices.Max(arr)+1)
	for _, x := range arr {
		cnt[x]++
	}
	slices.SortFunc(cnt, func(a, b int) int { return b - a })

	s := 0
	for i, c := range cnt {
		s += c
		if s >= len(arr)/2 {
			return i + 1
		}
	}
	return 0
}
