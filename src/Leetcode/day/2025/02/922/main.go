package main

/*
	给定一个非负整数数组 nums，  nums 中一半整数是 奇数 ，一半整数是 偶数 。

	对数组进行排序，以便当 nums[i] 为奇数时，i 也是 奇数 ；当 nums[i] 为偶数时， i 也是 偶数 。

	你可以返回 任何满足上述条件的数组作为答案 。
*/

func sortArrayByParityII(nums []int) []int {
	i, j := 0, 1
	for i < len(nums) {
		if nums[i]%2 == 0 {
			i += 2 // 寻找偶数下表最左边的奇数
		} else if nums[j]%2 == 1 {
			j += 2 // 寻找奇数下标中最左边的偶数
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i += 2
			j += 2
		}
	}
	return nums
}
