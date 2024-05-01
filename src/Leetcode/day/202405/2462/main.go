package main

/*
	给你一个下标从 0 开始的整数数组 costs ，其中 costs[i] 是雇佣第 i 位工人的代价。
	同时给你两个整数 k 和 candidates 。我们想根据以下规则恰好雇佣 k 位工人：

	· 总共进行 k 轮雇佣，且每一轮恰好雇佣一位工人。
	· 在每一轮雇佣中，从最前面 candidates 和最后面 candidates 人中选出代价最小的一位工人，如果有多位代价相同且最小的工人，选择下标更小的一位工人。
		· 比方说，costs = [3,2,7,7,1,2] 且 candidates = 2 ，第一轮雇佣中，我们选择第 4 位工人，因为他的代价最小 [3,2,7,7,1,2] 。
		· 第二轮雇佣，我们选择第 1 位工人，因为他们的代价与第 4 位工人一样都是最小代价，而且下标更小，[3,2,7,7,2] 。注意每一轮雇佣后，剩余工人的下标可能会发生变化。
	· 如果剩余员工数目不足 candidates 人，那么下一轮雇佣他们中代价最小的一人，如果有多位代价相同且最小的工人，选择下标更小的一位工人。
	· 一位工人只能被选择一次。
	返回雇佣恰好 k 位工人的总代价。

	雇佣过程可以用两个最小堆来模拟，一个负责维护 costs 剩余数字的最前面 candidates 个数的最小值，
							 另一个负责维护 costs 剩余数字的最后面 candidates 个数的最小值。

	1. 设 cost 的长度为 n。如果 candidates*2+k > n，我们一定可以选到 cost 中最小的 k 个数，所以直接返回 cost 的前 k 小之和。
	2. 初始化答案 ans = 0。初始化最小堆 pre 为 costs 最前面的 candidates 个数，初始化最小堆 suf 为 costs 最后面的 candidates 个数。
	   初始化坐标 i = candidates，j = n- 1 - candidates。
	3. 循环 k 次。每次循环，如果 pre 的堆顶小于 suf 的堆顶，则弹出 pre 的堆顶，加入答案，然后把 costs[i] 加入 pre，i 增加 1；
	   如果 suf 的堆顶小于 pre 的堆顶，则弹出 suf 的堆顶，加入答案，然后把 costs[j] 加入 suf，j 减少 1.
	4. 返回答案。
*/

import (
	"container/heap"
	"slices"
	"sort"
)

func totalCost(costs []int, k int, candidates int) (ans int64) {
	n := len(costs)
	if candidates*2+k > n {
		slices.Sort(costs)
		for _, x := range costs[:k] {
			ans += int64(x)
		}
		return
	}

	pre := hp{costs[:candidates]}
	suf := hp{costs[len(costs)-candidates:]}
	heap.Init(&pre)
	heap.Init(&suf)
	for i, j := candidates, n-1-candidates; k > 0; k-- {
		if pre.IntSlice[0] <= suf.IntSlice[0] {
			ans += int64(pre.replace(costs[i]))
			i++
		} else {
			ans += int64(suf.replace(costs[j]))
			j--
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(h, 0); return top }
