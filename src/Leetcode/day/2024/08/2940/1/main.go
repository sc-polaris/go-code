package main

import "container/heap"

/*
	给你一个下标从 0 开始的正整数数组 heights ，其中 heights[i] 表示第 i 栋建筑的高度。

	如果一个人在建筑 i ，且存在 i < j 的建筑 j 满足 heights[i] < heights[j] ，那么这个人可以移动到建筑 j 。

	给你另外一个数组 queries ，其中 queries[i] = [ai, bi] 。第 i 个查询中，Alice 在建筑 ai ，Bob 在建筑 bi 。

	请你能返回一个数组 ans ，其中 ans[i] 是第 i 个查询中，Alice 和 Bob 可以相遇的 最左边的建筑 。如果对于查询 i ，Alice 和 Bob 不能相遇，令 ans[i] 为 -1 。
*/

/*
	方法一：离线+最小堆
	离线：按照自己定义的某种顺序回答询问，而不是按照输入顺序 queries[0],queries[1],⋯ 回答询问。

	下文把 ai 和 bi 简称为 a 和 b。
	不妨设 a ≤ b。

	首先遍历 queries。如果 a = b 或者 heights[a] < heights[b]，那么 Alice 可以直接跳到 Bob 的位置，即 ans[i] = b。

	否则 heights[a] ≥ heights[b]，我们可以在位置 b 记录「左边有个 heights[a]，它属于第 i 个询问」，把数对（heights[a],i）加到列表 qs[b] 中。

	然后遍历 heights，同时用一个最小堆维护上面说的记录：遍历到 heights[i] 时，把 qs[i] 中的数对全部加入最小堆中。

	在加到最小堆前，我们可以回答堆中所有满足 heights[a] < heights[i] 的询问，由于 heights[b] ≤ heights[a] < heights[i]，所以该询问的答案是 i。

	为什么用最小堆？
	如果堆顶的 heights[a] ≥ heights[i]，那么堆中的其余元素也满足 heights[a] ≥ heights[i]，这些询问的答案肯定不是 i。

	总结
	算法涉及到三个位置，假定 a ≤ b，按照从左到右的顺序，它们分别是：
	1. a：回答询问时，用其高度 heights[a] 和当前高度 heights[i] 比大小，如果 heights[a] < heights[i] 则找到答案。
	2. b：决定了在什么位置把询问加入堆中。注意在遍历到位置 b 之前是不能入堆的。在遍历到位置 b 时入堆，这样后续只需要比较 heights[a] < heights[i]，如果
	   成立，就间接地说明 heights[b] < heights[i] 也成立。并且，由于我们是从左往右遍历 heights 的，当前下标 i 就是 Alice 和 Bob 可以相遇的最左边建筑的下标。
	3. 回答询问的位置 i。如果堆顶 heights[a] 小于当前位置的高度 heights[i]，则回答堆顶询问，并弹出堆顶。

*/

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}
	qs := make([][]pair, len(heights))
	for i, q := range queries {
		a, b := q[0], q[1]
		if a > b {
			a, b = b, a // 保证 a <= b
		}
		if a == b || heights[a] < heights[b] {
			ans[i] = b // a 直接跳到 b
		} else {
			qs[b] = append(qs[b], pair{heights[a], i}) // 离线询问
		}
	}

	h := hp{}
	for i, x := range heights {
		for h.Len() > 0 && h[0].h < x {
			// 堆顶的 heights[a] 可以跳到 heights[i]
			ans[heap.Pop(&h).(pair).i] = i
		}
		for _, p := range qs[i] {
			heap.Push(&h, p) // 后面再回答
		}
	}
	return ans
}

type pair struct{ h, i int }
type hp []pair

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].h < (*h)[j].h }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
