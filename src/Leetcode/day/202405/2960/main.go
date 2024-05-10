package main

/*
	给你一个长度为 n 、下标从 0 开始的整数数组 batteryPercentages ，表示 n 个设备的电池百分比。

	你的任务是按照顺序测试每个设备 i，执行以下测试操作：

	· 如果 batteryPercentages[i] 大于 0：
		· 增加 已测试设备的计数。
		· 将下标在 [i + 1, n - 1] 的所有设备的电池百分比减少 1，确保它们的电池百分比 不会低于 0 ，即 batteryPercentages[j] = max(0, batteryPercentages[j] - 1)。
		· 移动到下一个设备。
	· 否则，移动到下一个设备而不执行任何测试。
	返回一个整数，表示按顺序执行测试操作后 已测试设备 的数量。
*/

// 模拟

func countTestedDevices(batteryPercentages []int) (ans int) {
	for i := range batteryPercentages {
		if batteryPercentages[i] > 0 {
			ans++
			for j := i + 1; j < len(batteryPercentages); j++ {
				batteryPercentages[j] = max(0, batteryPercentages[j]-1)
			}
		}
	}
	return
}

/*
	1. 初始化 dev = 0，表示需要减 1 的次数
	2. 设 x = batteryPercentages[i]，那么该电池的实际百分比为 x - dec，如果 x - dec > 0，即
	   x > dec，那么后面的数都要减一，根据差分数组的思想，把 dec 加一即可。
	3. 答案就是 dec。因为每次遇到 x > dec 都把 dec 加一，这正是题目要求统计的。
*/

func countTestedDevices2(batteryPercentages []int) int {
	dec := 0
	for _, x := range batteryPercentages {
		if x > dec {
			dec++
		}
	}
	return dec
}
