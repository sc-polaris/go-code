package main

func carPooling(trips [][]int, capacity int) bool {
	d := [1001]int{}
	for _, t := range trips {
		d[t[1]] += t[0]
		d[t[2]] -= t[0]
	}
	s := 0
	for _, v := range d {
		s += v
		if s > capacity {
			return false
		}
	}
	return true
}
