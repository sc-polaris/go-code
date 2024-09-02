package main

func jump(nums []int) (ans int) {
	end, maxPos := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			end = maxPos
			ans++
		}
	}
	return ans
}
