package main

// 本质上是区间合并
func partitionLabels(s string) (ans []int) {
	var last [26]int
	for i, c := range s {
		last[c-'a'] = i // 每个字母最后出现的下标
	}

	start, end := 0, 0
	for i, c := range s {
		end = max(end, last[c-'a']) // 当前区间右端点的最大值
		if end == i {               // 区间合并完毕
			ans = append(ans, end-start+1)
			start = i + 1 // 下一个区间的左端点
		}
	}
	return
}
