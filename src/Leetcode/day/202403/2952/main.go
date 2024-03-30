package main

/*
	观察：
	为了方便描述，把 0 也算作可以得到的数。
	假设现在得到了区间 [0,s-1] 中的所有的数，如果此时遍历到正数 x = coins[i]，那么把 [0,s-1] 中的
	每个整数都增加 x，我们就得到了区间 [x,s+x-1] 中所有的数
	思路：
	把 coins 从小到大排序，遍历 x = coins[i]。分类讨论，看是否要添加数字：
	1. 如果 x <= s，那么，那么合并 [0,s-1] 和 [x,s+x-1] 这两个区间，我们可以得到 [0,s+x-1] 中的
	所有整数
	2. 如果 x >  s，或者遍历完了 coins 数组，意味着我们无法得到 s，那么就一定要把 s 加到 数组中（加
	一个比 s 还小的数字就没法得到更大的数，不够贪），这样就可以得到了 [s,2s-1] 中的所有整数，再与 [0,s-1]
	合并，可与你得到 [0,2s-1] 中的所有整数。然后再考虑 x 和 2s 的大小关系，继续分类讨论。

	但 s > target 时，我们就得到了 [1,target] 中的所有整数，退出循环。

*/

import "slices"

func minimumAddedCoins(coins []int, target int) (ans int) {
	slices.Sort(coins)
	s, i := 1, 0
	for s <= target {
		if i < len(coins) && coins[i] <= s {
			s += coins[i]
			i++
		} else {
			s *= 2 // 必须加 s
			ans++
		}
	}
	return
}
