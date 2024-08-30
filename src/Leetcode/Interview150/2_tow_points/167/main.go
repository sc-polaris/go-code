package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for {
		s := numbers[l] + numbers[r]
		if s == target {
			return []int{l + 1, r + 1}
		}
		if s > target {
			r--
		} else {
			l++
		}
	}
}
