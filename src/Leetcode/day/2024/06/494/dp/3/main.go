package main

import "math"

func findTargetSumWays(nums []int, target int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	s -= int(math.Abs(float64(target)))
	if s < 0 || s%2 == 1 {
		return 0
	}
	m := s / 2

	f := make([]int, m+1)
	f[0] = 1
	for _, x := range nums {
		for c := m; c >= x; c-- {
			f[c] += f[c-x]
		}
	}
	return f[m]
}
