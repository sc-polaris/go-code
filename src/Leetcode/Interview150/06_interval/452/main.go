package main

import "slices"

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int { return a[1] - b[1] })
	mr := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > mr {
			mr = p[1]
			ans++
		}
	}
	return ans
}
