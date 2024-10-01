package main

import (
	"container/heap"
	"slices"
	"sort"
)

/*
	有 n 名工人。 给定两个数组 quality 和 wage ，其中，quality[i] 表示第 i 名工人的工作质量，其最低期望工资为 wage[i] 。

	现在我们想雇佣 k 名工人组成一个工资组。在雇佣 一组 k 名工人时，我们必须按照下述规则向他们支付工资：

	1. 对工资组中的每名工人，应当按其工作质量与同组其他工人的工作质量的比例来支付工资。
	2. 工资组中的每名工人至少应当得到他们的最低期望工资。
	给定整数 k，返回 “组成满足上述条件的付费群体所需的最小金额”。在实际答案的 10^-5 以内的答案将被接受。

	tips：
	1. 在最优发工资的方案下，至少有一名工人，发给他的工资恰好等于他的最低期望工资。
	2. 枚举发了最低期望工资的那名工人，在满足题干中规则 1 的前提下，哪些共恩可以满足规则 2 ？如何快速地求出这些工人？
	3. 定义 r[i] = wage[i] / quality[i]，表示「每单位工作质量的工资」。
	   若以某人的 r[i] 为基准发工资，那么对于 r 值不超过 r[i] 的工人，发给他们的工资是不低于其最低期望工资的，因此这些
	   工人是可以随意选择（雇佣）的。
	4. 设这 k 名工人的 quality 之和为 sumQ，若以 r[i] 为基准发工资，那么发的总工资为 sumQ * r[i]，
	   因为 sumQ 越小发的工资总额就越小。
	   因此，我们需要在从小到大枚举 r[i] 时，维护当前最小的 k 个 quality 值。
	5. 用一个 “最大堆” 来维护。
	   按照 r[i] 从小到大顺序遍历工人，当堆中有 k 个元素时，如果 quality[i] 比堆顶小，则可以弹出顶，将 quality[i] 入堆，
	   从而得到一个更小的 sumQ，此时有可能找到一个更优解 sumQ * r[i]，更新答案。
*/

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	type pair struct{ q, w int }
	pairs := make([]pair, len(quality))
	for i, q := range quality {
		pairs[i] = pair{q, wage[i]}
	}
	// a.w/a.q - b.w/b.q = a.w*b.q - b.w*a.q
	slices.SortFunc(pairs, func(a, b pair) int { return a.w*b.q - b.w*a.q })

	h := hp{make([]int, k)}
	sumQ := 0
	for i, p := range pairs[:k] {
		h.IntSlice[i] = p.q
		sumQ += p.q
	}
	heap.Init(&h)

	ans := float64(sumQ*pairs[k-1].w) / float64(pairs[k-1].q) // 选 r 值最小的 k 名工人

	for _, p := range pairs[k:] { // 后面的工人 r 更大
		if p.q < h.IntSlice[0] { // 但是 sumQ 可以变小，从而可能得到更优的档案
			sumQ -= h.IntSlice[0] - p.q
			h.IntSlice[0] = p.q
			heap.Fix(&h, 0) // 更新堆顶
			ans = min(ans, float64(sumQ*p.w)/float64(p.q))
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *hp) Push(any)           {}
func (h *hp) Pop() (_ any)       { return }
