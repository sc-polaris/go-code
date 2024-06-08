package main

/*
	定义 f[i][j] 表示当前剩余元素从 nums[i] 到 nums[j]，此时最多可以进行的操作次数

	为避免出现 j = -1 的状态，需要把 f[i][j] 中的 j 加一（相当于再最左边插入一列），即：
	f[i][j+1] 表示当前剩余元素从 nums[i] 到 nums[j]，此时最多可以进行的操作次数

	注：如果记忆化搜索中的三个 if 都不成立，就不会继续递归，但递推需要计算所有状态。在随机数据下，本题递推效率不如记忆化搜索。
*/

func maxOperations(nums []int) int {
	n := len(nums)
	res1 := helper(nums[2:n], nums[0]+nums[1])      // 删除前两个数
	res2 := helper(nums[:n-2], nums[n-2]+nums[n-1]) // 删除后两个数
	res3 := helper(nums[1:n-1], nums[0]+nums[n-1])  // 删除第一个和最后一个数
	return max(res1, res2, res3) + 1                // 加上第一次操作
}

func helper(a []int, target int) int {
	n := len(a)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if a[i]+a[i+1] == target { // 删除前两个数
				f[i][j+1] = max(f[i][j+1], f[i+2][j+1]+1)
			}
			if a[j-1]+a[j] == target { // 删除后两个数
				f[i][j+1] = max(f[i][j+1], f[i][j-1]+1)
			}
			if a[i]+a[j] == target { // 删除第一个和最后一个数
				f[i][j+1] = max(f[i][j+1], f[i+1][j]+1)
			}
		}
	}
	return f[0][n]
}
