package main

import "slices"

/*
	首先从后向前查找第一个顺序对(i,i+1)，满足a[i]<a[i+1]。这样「较小数」即为a[i]。此时[i+1,n)必然是下降序列。
	如果找到了顺序对，那么在区间[i+1,n)中从后向前查找第一个元素j满足a[i]<a[j]。这样「较大数」即为a[j]。
	交换a[i]与a[j]，此时可以证明区间[i+1,n)必为降序。我们可以直接使用双指针反转区间[i+1,n)使其变为升序，而无需对该区间进行排序。
*/

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	slices.Reverse(nums[i+1:])
}
