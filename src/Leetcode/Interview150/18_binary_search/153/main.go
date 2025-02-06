package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= nums[len(nums)-1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
	//return nums[sort.Search(len(nums)-1, func(i int) bool { return nums[i] > nums[len(nums)-1] })]
}
