package main

/*
	时间复杂度的瓶颈在排序上，能否不排序呢？

	举例说明，如果 changed=[1,1,2,2,2,4,8,16,3,6,24,48,9,9,18,18]，我们把 x 和 2x 都放到同一组内，则
	可以分为如下四组：
	· 1,1,2,2,2,4,8,16
	· 3,6
	· 24,48
	· 9,9,18,18

	同一组内的数字，可以一并处理。
	1,1,2,2,2,4,8,16 这组数字，从 1 开始：
	· 用两个 1 消掉两个 2，把两个 1 加入答案。
	· 用剩下的一个 2 消掉一个 4，把一个 2 加入答案
	· 用一个 8 消掉一个 16，把一个 8 加入答案。
	其它组也类似。

	整理思路：
	1. 我们是从一组内的最小数字开始的。换句话说，只要 x/2 不在数组内，我们就可以从 x 开始。
	2. 从 x 开始，通过不断地乘 2，找到这一组内的其它数字。

	实现以上两点，完全不需要排序，算法如下：
	1. 用哈希表 cnt 统计 changed 数组的每个元素的出现次数。
	2. 遍历 cnt 中的元素。
	3. 设当前遍历到元素 x。如果 x 是偶数且 x/2 在 cnt 中，则跳过。
	4. 否则，我们可以从 x 开始循环，把 x,2x,4x,8x,... 全部消除（配对）。
	5. 每次循环，把 cnt[x] 个 x 和 cnt[x] 个 2x 配对。如果 x 的个数比 2x 的个数还多（或者 2x 不在 cnt 中），则无法配对，
	   返回空数组。否则把 cnt[x] 个 x 加入答案。然后分类讨论：
	   · 如果 cnt[x] < cnt[2x]，那么 cnt[2x] 减少 cnt[x]，把 x 乘以 2，继续循环。
	   · 如果 cnt[x] = cnt[2x]，说明 2x 也配对完了，把 x 乘以 4，继续循环。
	6. 最后返回答案。

	注意 0 要单独处理，因为一个 0 可以和其它 0 消除。这意味着 changed 中必须要有偶数个 0。设 changed 中有 cnt0 个 0，如果
	cnt0 是奇数，返回空数组，否则把 cnt0/2 个 0 加入答案。
*/

func findOriginalArray(changed []int) []int {
	cnt := make(map[int]int)
	for _, x := range changed {
		cnt[x]++
	}

	cnt0 := cnt[0]
	if cnt0%2 == 1 {
		return nil
	}
	delete(cnt, 0)
	ans := make([]int, cnt0/2, len(changed)/2)

	for x := range cnt {
		// 如果 x/2 在 cnt 中，则跳过
		if x%2 == 0 && cnt[x/2] > 0 {
			continue
		}
		// 把 x, 2x, 4x, 8x, ... 全都配对
		for cnt[x] > 0 {
			// 每次循环，把 cntX 个 x 和 cntX 个 2x 配对
			cntX := cnt[x]
			// 无法配对，至少需要有 cntX 个 2x
			if cntX > cnt[x*2] {
				return nil
			}
			for range cntX {
				ans = append(ans, x)
			}
			if cntX < cnt[x*2] {
				// 还剩下一些 2x
				cnt[x*2] -= cntX
				x *= 2
			} else {
				x *= 4
			}
		}
	}
	return ans
}
