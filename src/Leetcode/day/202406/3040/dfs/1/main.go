package main

/*
	给你一个整数数组 nums ，如果 nums 至少 包含 2 个元素，你可以执行以下操作中的 任意 一个：
	· 选择 nums 中最前面两个元素并且删除它们。
	· 选择 nums 中最后两个元素并且删除它们。
	· 选择 nums 中第一个和最后一个元素并且删除它们。
	一次操作的 分数 是被删除元素的和。

	在确保 所有操作分数相同 的前提下，请你求出 最多 能进行多少次操作。

	请你返回按照上述要求 最多 可以进行的操作次数。
*/

func maxOperations(nums []int) int {
	n := len(nums)
	res1 := help(nums[2:], nums[0]+nums[1])       // 删除前两个数
	res2 := help(nums[:n-2], nums[n-2]+nums[n-1]) // 删除后两个数
	res3 := help(nums[1:n-1], nums[0]+nums[n-1])  // 删除第一个和最后一个数
	return max(res1, res2, res3) + 1              // 加上第一次操作
}

func help(a []int, target int) int {
	n := len(a)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i >= j {
			return
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		if a[i]+a[i+1] == target { // 删除前两个数
			res = max(res, dfs(i+2, j)+1)
		}
		if a[j-1]+a[j] == target { // 删除后两个数
			res = max(res, dfs(i, j-2)+1)
		}
		if a[i]+a[j] == target { // 删除前后两个数
			res = max(res, dfs(i+1, j-1)+1)
		}
		*p = res // 记忆化
		return
	}
	return dfs(0, n-1)
}
