package main

/*
	给你一个下标从 0 开始的二维整数数组 nums 表示汽车停放在数轴上的坐标。对于任意下标 i，nums[i] = [starti, endi] ，
	其中 starti 是第 i 辆车的起点，endi 是第 i 辆车的终点。

	返回数轴上被车 任意部分 覆盖的整数点的数目。
*/

/*
	核心思路：计算每个点被覆盖了多少次。统计覆盖次数大于 0 的点，即为答案。
	假设一开始有一个全为 0 的数组 a，用来保存每个点被覆盖了多少次。
	对于示例 1，我们可以把 a 中下标在 [3,6] 的元素都加一，下标在 [1,5] 的元素都加一，下标在 [4,7] 的元素都加一。
	然后，统计 a[i]>0 的个数，即为答案。

	如何快速地「把区间内的数都加一」呢？
	这可以用差分数组实现。

	如果我们维护数组 diff，那么 count[i] 可以通过从 diff[0] 累加到 diff[i] 方便地求出。
	当我们需要将数组 count 中下标从 x 到 y 的元素均增加 1 时，对应到数组 diff，只需要将 diff[x] 增加 1，并将 diff[y+1] 减少 1，时间复杂度从 O(C) 降低至 O(1)。

	最后只需要对数组 diff 求一遍前缀和，就还原出了数组 count，其中非零元素的数量即为答案。
*/

func numberOfPoints(nums [][]int) (ans int) {
	maxEnd := 0
	for _, interval := range nums {
		maxEnd = max(maxEnd, interval[1])
	}

	diff := make([]int, maxEnd+2) // 注意下面有 end+1
	for _, interval := range nums {
		diff[interval[0]]++
		diff[interval[1]+1]--
	}

	s := 0
	for _, d := range diff {
		s += d
		if s > 0 {
			ans++
		}
	}
	return
}
