package main

/*
	定理：如果 i<n−1 且 nums[i]<nums[i+1]，那么在下标 [i+1,n−1] 中一定存在至少一个峰值。
	证明：反证法，假设下标 [i+1,n−1] 中没有峰值。
		· 由于 i+1 不是峰值且 nums[i]<nums[i+1]，所以一定有 nums[i+1]<nums[i+2] 成立，否则 i+1 就是峰值了。注意题目保证相邻元素不同，不存在相邻元素相等的情况。
		· 由于 i+2 不是峰值且 nums[i+1]<nums[i+2]，所以一定有 nums[i+2]<nums[i+3] 成立，否则 i+2 就是峰值了。
	依此类推，得
			nums[i]<nums[i+1]<nums[i+2]<⋯<nums[n−1]>nums[n]=−∞
	这意味着 nums[n−1] 是峰值，矛盾，所以原命题成立。
	同理可得，如果 i<n−1 且 nums[i]>nums[i+1]，那么在 [0,i] 中一定存在至少一个峰值。

	所以，通过比较 nums[i] 和 nums[i+1] 的大小关系，从而不断地缩小峰值所在位置的范围，二分找到峰值。
*/

func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] <= nums[mid+1] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
	//return sort.Search(len(nums)-1, func(i int) bool { return nums[i] > nums[i+1] })
}
