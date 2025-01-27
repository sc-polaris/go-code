package main

/*
	给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

	每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
	· 0 <= j <= nums[i]
	· i + j < n
	返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
*/

func jump(nums []int) (ans int) {
	curRight := 0  // 已建造的桥的右端点
	nextRight := 0 // 下一座桥的右端点的最大值
	for i, num := range nums[:len(nums)-1] {
		nextRight = max(nextRight, i+num)
		if i == curRight { // 到达已建造的桥的右端点
			curRight = nextRight // 造一座桥
			ans++
		}
	}
	return
}
