package main

/*
	给你一个整数数组 nums ，一个整数数组 queries 和一个整数 x 。

	对于每个查询 queries[i] ，你需要找到 nums 中第 queries[i] 个 x 的位置，并返回它的下标。如果数组中 x 的出现次数少于 queries[i] ，该查询的答案为 -1 。

	请你返回一个整数数组 answer ，包含所有查询的答案。
*/

func occurrencesOfElement(nums []int, queries []int, x int) (ans []int) {
	idx := 0
	pos := make([]int, len(nums))
	for i, v := range nums {
		if v == x {
			idx++
			pos[idx] = i
		}
	}
	for _, v := range queries {
		if v > idx {
			ans = append(ans, -1)
		} else {
			ans = append(ans, pos[v])
		}
	}
	return
}

// 灵神写法

func occurrencesOfElement2(nums []int, queries []int, x int) (ans []int) {
	var pos []int
	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}
	for i, q := range queries {
		if q > len(pos) {
			queries[i] = -1
		} else {
			queries[i] = pos[q-1]
		}
	}
	return queries
}
