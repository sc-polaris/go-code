package main

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		res += sort.SearchInts(nums[0:i], target-nums[i])
	}
	return res
}

func countPairs2(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	for i, j := 0, len(nums)-1; i < j; i++ {
		for i < j && nums[i]+nums[j] >= target {
			j--
		}
		res += j - i
	}
	return res
}

func main() {

}
