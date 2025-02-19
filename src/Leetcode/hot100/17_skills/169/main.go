package main

func majorityElement(nums []int) int {
	ans, cnt := 0, 0
	for _, x := range nums {
		if cnt == 0 {
			ans = x
		}
		if ans == x {
			cnt++
		} else {
			cnt--
		}
	}
	return ans
}
