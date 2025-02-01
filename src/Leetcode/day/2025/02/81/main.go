package main

/*
	已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。

	在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。

	给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。

	你必须尽可能减少整个操作步骤。
*/

func search(nums []int, target int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}

	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == target || nums[r] == target {
			return true
		}
		l++
		r--
	}
	return false
}

func search2(nums []int, target int) bool {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)

	check := func(i int) bool {
		x := nums[i]
		if x > nums[right] {
			return target > nums[right] && x >= target
		}
		return target > nums[right] || x >= target
	}

	for left+1 < right { // 开区间不为空
		mid := (left + right) / 2
		if nums[mid] == nums[right] {
			right--
		} else if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right] == target
}
