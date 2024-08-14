package main

/*
			0,	nums[i] mod 2 != nums[i+1] mod 2
	a[i] =
			1,	nums[i] mod 2 == nums[i+1] mod 2

	如果 a 的下标从 from 到 to−1 的子数组和等于 0，就说明 nums 的下标从 from 到 to 的这个子数组，其所有相邻元素的奇偶性都不同，该子数组为特殊数组。

	计算 a 的前缀和 s，可以快速判断子数组和是否为 0，也就是判断
				s[to]−s[from] = 0
	即
				s[from] = s[to]
*/

func isArraySpecial(nums []int, queries [][]int) []bool {
	s := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		s[i] = s[i-1]
		if nums[i-1]%2 == nums[i]%2 {
			s[i]++
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = s[q[0]] == s[q[1]]
	}
	return ans
}
