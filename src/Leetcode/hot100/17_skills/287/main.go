package main

/*
	二分查找除了对索引二分，还有值域二分
	数组元素是 1 - n 中的某一个，出现的位置不确定，但值域是确定的。
		对索引二分，一般用于有序数组中找元素，因为索引的大小可以反映值的大小，因此对索引二分即可。
		对值域二分。重复数落在 [1, n] ，可以对 [1, n] 这个值域二分查找。
	mid = (1 + n) / 2，重复数要么落在[1, mid]，要么落在[mid + 1, n]。
	遍历原数组，统计 <= mid 的元素个数，记为 k。
	如果k > mid，说明有超过 mid 个数落在[1, mid]，但该区间只有 mid 个“坑”，说明重复的数落在[1, mid]。
	相反，如果k <= mid，则说明重复数落在[mid + 1, n]。
	对重复数所在的区间继续二分，直到区间闭合，重复数就找到了。

*/

func findDuplicate(nums []int) int {
	l, r := 1, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		count := 0
		for _, x := range nums {
			if x <= mid {
				count++
			}
		}
		if count <= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func findDuplicate2(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
