package main

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
