package enum

/*
nums[l] % 2 == 0
对于范围 [l, r - 1] 内的所有下标 i ，nums[i] % 2 != nums[i + 1] % 2
对于范围 [l, r] 内的所有下标 i ，nums[i] <= threshold
*/

func check(nums []int, threshold int) bool {
	if nums[0]%2 != 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > threshold || (i+1 < len(nums)) && nums[i]%2 == nums[i+1]%2 {
			return false
		}
	}
	return true
}

func longestAlternatingSubarray(nums []int, threshold int) (res int) {
	for l := 0; l < len(nums); l++ {
		for r := l; r < len(nums); r++ {
			if check(nums[l:r+1], threshold) && res < r-l+1 {
				res = r - l + 1
			}
		}
	}
	return
}
