package main

/*
	给你一个长度为 n 的整数数组 nums 和一个正整数 k 。

	一个数组的 能量值 定义为：
	· 如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素。
	· 否则为 -1 。
	你需要求出 nums 中所有长度为 k 的子数组的能量值。

	请你返回一个长度为 n - k + 1 的整数数组 results ，其中 results[i] 是子数组 nums[i..(i + k - 1)] 的能量值。
*/

/*
	核心思路：找连续上升的段。如果段长至少是 k，那么这段中的所有长为 k 的子数组都是符合要求的，子数组的最后一个元素是最大的。
	具体来说，遍历数组的同时，用一个计数器 cnt 统计连续递增的元素个数：
	· 初始化 cnt=0。
	· 如果 i=0 或者 nums[i]=nums[i−1]+1，把 cnt 增加 1；否则，把 cnt 置为 1。
	· 如果发现 cnt≥k，那么下标从 i−k+1 到 i 的这个子数组的能量值为 nums[i]，即 ans[i−k+1]=nums[i]。
*/

func resultsArray(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	for i := range ans {
		ans[i] = -1
	}
	cnt := 0
	for i, x := range nums {
		if i == 0 || x == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			ans[i-k+1] = x
		}
	}
	return ans
}
