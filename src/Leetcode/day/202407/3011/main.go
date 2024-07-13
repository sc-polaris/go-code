package main

import (
	"math/bits"
	"slices"
)

/*
	给你一个下标从 0 开始且全是 正 整数的数组 nums 。

	一次 操作 中，如果两个 相邻 元素在二进制下数位为 1 的数目 相同 ，那么你可以将这两个元素交换。你可以执行这个操作 任意次 （也可以 0 次）。

	如果你可以使数组变有序，请你返回 true ，否则返回 false 。
*/

/*
	方法一：直接排序

	方法二：记录每一段的最小值和最大值
		对于每一组，如果这一组的每个数，都大于等于上一组的最大值 preMax，那么我们就能把数组排成递增的，否则不行。
		由于题目保证 nums[i]>0，我们可以把 preMax 和本组最大值 mx 都初始化成 0。
*/

func canSortArray(nums []int) bool {
	for i, n := 0, len(nums); i < n; {
		start := i
		ones := bits.OnesCount(uint(nums[i]))
		i++
		for i < n && bits.OnesCount(uint(nums[i])) == ones {
			i++
		}
		slices.Sort(nums[start:i])
	}
	return slices.IsSorted(nums)
}

func canSortArray2(nums []int) bool {
	preMax := 0
	for i, n := 0, len(nums); i < n; {
		mx := 0
		ones := bits.OnesCount(uint(nums[i]))
		for ; i < n && bits.OnesCount(uint(nums[i])) == ones; i++ {
			if nums[i] < preMax {
				return false
			}
			mx = max(mx, nums[i])
		}
		preMax = mx
	}
	return true
}
