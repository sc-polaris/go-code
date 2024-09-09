package main

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	l, r := newInterval[0], newInterval[1]
	merged := false
	for _, p := range intervals {
		if p[0] > r {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{l, r})
				merged = true
			}
			ans = append(ans, p)
		} else if p[1] < l {
			// 在插入区间的左侧且无交集
			ans = append(ans, p)
		} else {
			// 与插入区间有交集，计算它们的并集
			l = min(l, p[0])
			r = max(r, p[1])
		}
	}
	if !merged {
		ans = append(ans, []int{l, r})
	}
	return
}
