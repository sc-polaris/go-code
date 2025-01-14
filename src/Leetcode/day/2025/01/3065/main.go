package main

/*
	给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

	一次操作中，你可以删除 nums 中的最小元素。

	你需要使数组中的所有元素都大于或等于 k ，请你返回需要的 最少 操作次数。
*/

func minOperations(nums []int, k int) (ans int) {
	for _, x := range nums {
		if x < k {
			ans++
		}
	}
	return
}
