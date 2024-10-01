package main

import "math"

/*
	给你一个下标从 0 开始的二进制数组 nums，其长度为 n ；另给你一个 正整数 k 以及一个 非负整数 maxChanges 。

	Alice 在玩一个游戏，游戏的目标是让 Alice 使用 最少 数量的 行动 次数从 nums 中拾起 k 个 1 。游戏开始时，Alice
	可以选择数组 [0, n - 1] 范围内的任何索引 aliceIndex 站立。如果 nums[aliceIndex] == 1 ，Alice 会拾起一个 1 ，
	并且 nums[aliceIndex] 变成0（这 不算 作一次行动）。之后，Alice 可以执行 任意数量 的 行动（包括零次），在每次行动
	中 Alice 必须 恰好 执行以下动作之一：
	· 选择任意一个下标 j != aliceIndex 且满足 nums[j] == 0 ，然后将 nums[j] 设置为 1 。这个动作最多可以执行 maxChanges 次。
	· 选择任意两个相邻的下标 x 和 y（|x - y| == 1）且满足 nums[x] == 1, nums[y] == 0 ，然后交换它们的值（将 nums[y] = 1 和
	  nums[x] = 0）。如果 y == aliceIndex，在这次行动后 Alice 拾起一个 1 ，并且 nums[y] 变成 0 。
	返回 Alice 拾起 恰好 k 个 1 所需的 最少 行动次数。

	贪心
*/

/*
	把 0 看作「空位」
	第二种操作相当故意把一个 1 移动到和它相邻的空位三，如果我们想得到一个下标在 j 的 1，必须操作 ∣aliceIndex−j∣ 次。

	对于第一种操作，贪心地把和 aliceIndex 相邻的 0 变成 1（在此之前先移动相邻的 1），然后结合第二种操作，把相邻的 1
	移动到 aliceIndex，只需 2 次操作就可以得到一个 1。

	我们分 maxChanges 较大，和 maxChanges 较小两种情况讨论。

	maxChanges 较大的情况
	应当优先使用第一种操作+第二种操作，毕竟只需要操作 2 次就能得到一个 1。那么答案就是 2k 吗？
	细节：对于 aliceIndex,aliceIndex−1,aliceIndex+1 这三个位置上的 1，可以用更少的操作得到：
	· aliceIndex 位置上的 1 无需操作就能得到
	· aliceIndex−1 和 aliceIndex+1 位置上的 1 只需操作 1 次就能得到。

	贪心的想法是，选择有三个连续 1 的中间位置，作为 aliceIndex。如果没有三个连续 1，就看有没有连续两个 1。如果没有连
	续两个 1，就选任意 1 的位置。如果没有 1 就随便选。

	一般地，设 c 为 nums 中的长度不超过 3 的最长连续 1 的长度。如果 c>k 则 c=k。

	如果 maxChanges≥k−c，我们可以先使用 max(c−1,0) 次第二种操作，收集这连续的 c 个 1，然后对于其余 k−c 个 1，都可
	以用 2 次操作得到，此时可以直接返回 max(c−1,0)+(k−c)⋅2。

	接下来，要解决的就是 maxChanges 比较小的情况了。
	从特殊到一般，想一想，如果 maxChanges=0，也就是只能使用第二种操作，要如何计算答案呢？

	maxChanges=0 的情况
	首先算出所有 1 的位置，记到一个 pos 数组中。例如示例 1 的 nums=[1,1,0,0,0,1,1,0,0,1]，其 pos=[0,1,5,6,9]。
	示例 1 的 k=3，我们可以枚举 pos 的所有长为 3 的子数组，例如 [0,1,5]，就好比在坐标轴上的 0,1,5 位置上有 3 个生产商
	品的工厂，我们要建造一个货仓存放商品，把货仓建在哪里，可以使所有工厂到货仓的距离之和最小？

	这个问题叫做「货仓选址」。根据 中位数贪心及其证明，最优解是把货仓建在工厂位置的中位数上。例如 [0,1,5] 中的 1，此时距离
	和等于 ∣0−1∣+∣1−1∣+∣5−1∣=5。

	maxChanges 较小的情况
	最后，如果 maxChanges>0，我们可以先计算所有长为 k−maxChanges 的子数组的货仓选址问题，取最小值，然后再通过 maxChanges⋅2 次操作得到 maxChanges 个 1。

	示例 1 只需考虑所有长为 k−1=2 的子数组，那么前两个 1 的货仓选址问题就是最小的，距离之和为 1，也就是这两个 1 需要 1 次
	操作得到。然后再通过 2 次操作得到剩下的一个 1，总共需要 1+2=3 次操作。
*/

func minimumMoves(nums []int, k int, maxChanges int) int64 {
	var pos []int
	c := 0 // nums 中连续的 1 长度
	for i, x := range nums {
		if x == 0 {
			continue
		}
		pos = append(pos, i) // 记录 1 的位置
		c = max(c, 1)
		if i > 0 && nums[i-1] == 1 {
			if i > 1 && nums[i-2] == 1 {
				c = 3 // 有 3 个 连续的 1
			} else {
				c = max(c, 2) // 有 2 个连续的 1
			}
		}
	}

	c = min(c, k)
	if maxChanges >= k-c {
		// 其余 k-c 个 1 可以全部用两次操作得到
		return int64(max(c-1, 0) + (k-c)*2)
	}

	n := len(pos)
	sum := make([]int, n+1)
	for i, x := range pos {
		sum[i+1] = sum[i] + x
	}

	ans := math.MaxInt
	// 除了 maxChanges 个数可以用两次操作得到，其余的 1 只能一步步移动到 pos[i]
	size := k - maxChanges
	for right := size; right <= n; right++ {
		// s1+s2 是 j 在 [left, right) 中的所有 pos[j] 到 pos[(left+right)/2] 的距离之和
		left := right - size
		i := left + size/2
		s1 := pos[i]*(i-left) - (sum[i] - sum[left])
		s2 := sum[right] - sum[i] - pos[i]*(right-i)
		ans = min(ans, s1+s2)
	}
	return int64(ans + maxChanges*2)
}
