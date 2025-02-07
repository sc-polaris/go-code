package main

func singleNumber(nums []int) (res int) {
	for _, x := range nums {
		res ^= x
	}
	return
}
