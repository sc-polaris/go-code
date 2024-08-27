package main

import "sort"

/*
	给你一个整数数组 nums 。数组 nums 的 唯一性数组 是一个按元素从小到大排序的数组，包含了 nums 的所有
	非空子数组中
	不同元素的个数。

	换句话说，这是由所有 0 <= i <= j < nums.length 的 distinct(nums[i..j]) 组成的递增数组。

	其中，distinct(nums[i..j]) 表示从下标 i 到下标 j 的子数组中不同元素的数量。

	返回 nums 唯一性数组 的 中位数 。

	注意，数组的 中位数 定义为有序数组的中间元素。如果有两个中间元素，则取值较小的那个。
*/

/*
	提示 1：二分答案
	nums 的唯一性数组有多少个？也就是说 nums 的非空连续子数组的个数。
	长为 n 的子数组有 1 个，长为 n-1 的子数组有 2 个，......，长为 1 的子数组有 n 个。
	所以一共有 m = 1 + 2 + ... + n = n(n+1)/2 个非空连续子数组。

	这 m 个子数组，对应着 m 个 distinct 值。
	中位数是这 m 个数中的第 k = ⌈ m/2 ⌉。例如 m = 4 时，中位数是其中第 2 小元素。
	考虑这 m 个数中，小于等于某个定值 upper 的数有多少个？
	由于 upper 越大，小于等于 upper 的数越多，有单调性，故可以二分中位数为 upper，问题变成：
	· distinct 值 <= upper 的子数组有多少个？
	设子数组的个数为 cnt，如果 cnt < k 说明二分的 upper 小了，更新二分左边界 left，否则更新二分右边界 right。

	提示 2：滑动窗口
	怎么计算 distinct 值 <= upper 的子数组个数？
	由于子数组越长，不同元素个数（distinct 值）不会变小，有单调性，故可以用滑动窗口计算子数组个数。
	用一个哈希表 freq 统计窗口（子数组）内的元素及其出现次数。

	枚举窗口右端点 r，把 nums[r] 加入 freq（出现次数加一）。如果发现 freq 的大小超过 upper，说明窗口内的元素过多，那么
	不断移出窗口左端点元素 nums[l]（出现次数减一，如果出现次数等于 0 就从 freq 中移除），直到 freq 的大小 <= upper 为止。

	此时右端点为 r，左端点为 l,l+1,l+2,...,r 的子数组都是满足要求的（distinct 值 <= upper），一共有 r-l+1 个，加到
	子数组个数 cnt 中。
*/

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	k := (n*(n+1)/2 + 1) / 2
	ans := sort.Search(n, func(upper int) bool {
		cnt := 0
		l := 0
		freq := make(map[int]int)
		for r, in := range nums {
			freq[in]++              // 移入右端点
			for len(freq) > upper { // 窗口内元素过多
				out := nums[l]
				freq[nums[l]]-- // 移出左端点
				if freq[out] == 0 {
					delete(freq, out)
				}
				l++
			}
			cnt += r - l + 1 // 右端点固定为 r 时，有 r-l+1 个合法左端点
			if cnt >= k {
				return true
			}
		}
		return false
	})
	// 特判
	if ans == 0 {
		return 1
	}
	return ans
}
