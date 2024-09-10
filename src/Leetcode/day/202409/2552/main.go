package main

/*
	给你一个长度为 n 下标从 0 开始的整数数组 nums ，它包含 1 到 n 的所有数字，请你返回上升四元组的数目。

	如果一个四元组 (i, j, k, l) 满足以下条件，我们称它是上升的：
	1. 0 <= i < j < k < l < n 且
	2. nums[i] < nums[k] < nums[j] < nums[l] 。
*/

/*
	核心思想：
	枚举右边的 4（下标l），问题变成维护左边的 132 模式的个数，注意 132 模式的 3 必须比 nums[l] 小。
	如果发现 nums[j] < nums[l]，那么就把 [0,l-1] 中的 3 在下标 j 的 132 模式的个数加入答案。

	对于 132 模式，可以枚举中间的3（下标 j），问题变成维护 12 模式的个数。

	12 模式
	枚举 2（下标 k）和 1（下标 i）：
	· 如果 nums[i] < nums[k]，则找到了一个 2 在下标 k 的 12 模式。

	132 模式
	枚举 2（下标 k）和 3（下标 j）：
	· 定义 cnt3[j] 表示 3 的下标为 j 时的 132 模式的个数。
	· 定义 cnt2 表示 2 的下标为 k 时的 12 模式的个数。
	分类讨论
	· 如果 nums[j]>nums[k]，把 cnt2 个 12 模式加到 cnt3[j] 中。
	· 把 j 当作 i，如果 nums[i]<nums[k]，我们找到了一个 12 模式，把 cnt2 加一。
	注意：我们并不需要单独计算 12 模式的个数，而是把 12 模式的计算过程整合到 132 模式的计算过程中。

	1324 模式
	枚举 l 和 j，分类讨论：
	· 如果 nums[j]<nums[l]，把 cnt3[j] 个 132 模式加到答案中。
	· 把 j 当作 i，把 l 当作 k，如果 nums[i]<nums[k]，我们找到了一个 12 模式，把 cnt2 加一。
	· 把 l 当作 k，如果 nums[j]>nums[k]，把 cnt2 个 12 模式加到 cnt3[j] 中。
	注意：我们并不需要单独计算 132 模式的个数，而是把 132 模式的计算过程整合到 1324 模式的计算过程中。
*/

func countQuadruplets(nums []int) (cnt4 int64) {
	cnt3 := make([]int, len(nums))
	for l := 2; l < len(nums); l++ {
		cnt2 := 0
		for j := 0; j < l; j++ {
			if nums[j] < nums[l] { // 3 < 4
				cnt4 += int64(cnt3[j])
				// 把 j 当做 i，把 l 当做 k，现在 nums[i] < nums[k]，即 1 < 2
				cnt2++
			} else { // 把 l 当作 k，现在 nums[j] > nums[k]，即 3 > 2
				cnt3[j] += cnt2
			}
		}
	}
	return
}
