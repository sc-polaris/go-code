package main

import "slices"

/*
	给你一个下标从 0 开始长度为 n 的整数数组 buses ，其中 buses[i] 表示第 i 辆公交车的出发时间。同时给你一个下标从 0
	开始长度为 m 的整数数组 passengers ，其中 passengers[j] 表示第 j 位乘客的到达时间。所有公交车出发的时间互不相同，
	所有乘客到达的时间也互不相同。

	给你一个整数 capacity ，表示每辆公交车 最多 能容纳的乘客数目。

	每位乘客都会搭乘下一辆有座位的公交车。如果你在 y 时刻到达，公交在 x 时刻出发，满足 y <= x  且公交没有满，那么你可以搭
	乘这一辆公交。最早 到达的乘客优先上车。

	返回你可以搭乘公交车的最晚到达公交站时间。你 不能 跟别的乘客同时刻到达。

	注意：数组 buses 和 passengers 不一定是有序的。
*/

/*
	分类讨论：
	1. 如果最后一班公交还有空位：
		1.1 如果最后一班公交发车时，没有乘客到达公交站，我们可以在发车时到达公交站。极限操作
		1.2 如果发车时恰好有乘客到达公交站，由于题目要求不能跟别的乘客同时到达，我们可以顺着这位乘客往前找没人到达的时刻，
			在这个时刻「插队」。
	2. 如果最后一班公交没有空位：
		2.1 找最后一个上车的乘客 A，然后往前找没人到达的时刻，在这个时刻「插队」，把 A 挤下去。（可怜的 A）

	为什么可以插队？万一前面的乘客没有上车，我们不就也没法上车了吗？
	这是不会的，因为先来先上车，如果一个乘客上了车，那么他前面的乘客也肯定上了车。

	模拟乘客上车
	1. 为方便模拟，把 buses 和 passengers 都从小到大排序。
	2. 双指针遍历 buses 和 passengers。
	3. 对于 buses[i]，初始化 c=capacity。
	4. 不断循环，如果 c>0 且 passengers[j]≤buses[i]，那么第 j 位乘客可以上车，把 c 减一，j 加一。如果没法上车，
	   只能等下一班车。
	5. 双指针遍历结束后，j−1 就是最后一个上车的乘客。这里减一是因为第 j 位乘客上车后我们把 j 加一了。
*/

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) (ans int) {
	slices.Sort(buses)
	slices.Sort(passengers)

	// 模拟乘客上车
	j, c := 0, 0
	for _, t := range buses {
		for c = capacity; c > 0 && j < len(passengers) && passengers[j] <= t; c-- {
			j++
		}
	}

	// 插队
	if c > 0 {
		ans = buses[len(buses)-1] // 最后一班公交还有空位，在发车时到达
	} else {
		ans = passengers[j-1] // 上一个上车的乘客
	}
	for j--; j >= 0 && ans == passengers[j]; j-- { // 往前找没人到达的时刻
		ans--
	}
	return
}
