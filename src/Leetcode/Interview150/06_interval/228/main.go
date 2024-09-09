package main

import "strconv"

func summaryRanges(nums []int) (ans []string) {
	i, n := 0, len(nums)
	for i < n {
		start, v := i, strconv.Itoa(nums[i])
		for i < n-1 && nums[i+1]-nums[i] == 1 {
			i++
		}
		if start != i {
			ans = append(ans, v+"->"+strconv.Itoa(nums[i]))
		} else {
			ans = append(ans, v)
		}
		i++
	}
	return
}
