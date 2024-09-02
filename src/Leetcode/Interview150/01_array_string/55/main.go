package main

func canJump(nums []int) bool {
	mx := 0
	for i, jump := range nums {
		if i > mx {
			return false
		}
		mx = max(mx, i+jump)
		if mx >= len(nums)-1 {
			break
		}
	}
	return true
}
