package dp

/*
nums[l] % 2 == 0
对于范围 [l, r - 1] 内的所有下标 i ，nums[i] % 2 != nums[i + 1] % 2
对于范围 [l, r] 内的所有下标 i ，nums[i] <= threshold
*/

func longestAlternatingSubarray(nums []int, threshold int) (res int) {
	dp := 0
	for l := len(nums) - 1; l >= 0; l-- {
		if nums[l] > threshold {
			dp = 0
		} else if l == len(nums)-1 || nums[l]%2 != nums[l+1]%2 {
			dp++
		} else {
			dp = 1
		}
		if nums[l]%2 == 0 && dp > res {
			res = dp
		}
	}
	return res
}
