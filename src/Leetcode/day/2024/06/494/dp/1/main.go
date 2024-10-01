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

	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	f[0][0] = 1
	for i, x := range nums {
		for c := 0; c <= m; c++ {
			if c < x {
				f[i+1][c] = f[i][c] // 只能不选
			} else {
				f[i+1][c] = f[i][c] + f[i][c-x] // 不选 + 选
			}
		}
	}
	return f[n][m]
}
