package main

/*
	x0-y0    y0-z0
    -----  = -----
    x1-y1    y1-z1
*/

func maxPoints(points [][]int) int {
	n := len(points)
	ans := 1
	for i, x := range points {
		for j := i + 1; j < n; j++ {
			y := points[j]
			cnt := 2
			for k := j + 1; k < n; k++ {
				z := points[k]
				s1 := (x[0] - y[0]) * (y[1] - z[1])
				s2 := (x[1] - y[1]) * (y[0] - z[0])
				if s1 == s2 {
					cnt++
				}
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}
