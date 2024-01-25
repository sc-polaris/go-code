package main

import "math/bits"

func sumIndicesWithKSetBits(nums []int, k int) int {
	ans := 0
	for i, x := range nums {
		if bits.OnesCount(uint(i)) == k {
			ans += x
		}
	}
	return ans
}

func sumIndicesWithKSetBits2(nums []int, k int) int {
	ans := 0
	for i, x := range nums {
		if bitCount(i) == k {
			ans += x
		}
	}
	return ans
}

func bitCount(x int) int {
	cnt := 0
	for x != 0 {
		cnt += x % 2
		x /= 2
	}
	return cnt
}

func main() {
}
