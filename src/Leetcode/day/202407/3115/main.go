package main

/*
	给你一个整数数组 nums。
	返回两个（不一定不同的）质数在 nums 中 下标 的 最大距离。
*/

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

func maximumPrimeDifference(nums []int) int {
	i := 0
	for !isPrime(nums[i]) {
		i++
	}
	j := len(nums) - 1
	for !isPrime(nums[j]) {
		j--
	}
	return j - i
}
