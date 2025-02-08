package main

func mySqrt(x int) int {
	l, r := 1, x+1
	for l < r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1 // l 可能是比 sqrt(x) 大 1 的值，所以返回 l-1
}
