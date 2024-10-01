package main

/*
	给你一个 二进制数组 nums 。

	如果一个 子数组中 不存在 两个 相邻 元素的值 相同 的情况，我们称这样的子数组为 交替子数组 。

	返回数组 nums 中交替子数组的数量。
*/

func countAlternatingSubarrays(nums []int) (ans int64) {
	cnt := 0
	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			cnt = 1
		} else {
			cnt++
		}
		ans += int64(cnt) // 有 cnt 个以 i 为右端点的交替子数组
	}
	return
}
