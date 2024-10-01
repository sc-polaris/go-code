package main

import "math/bits"

/*
	方法三：在线+线段树二分
	在线：按照输入顺序 queries[0],queries[1],⋯ 一个一个地回答询问。

	问题相当于计算区间 [b+1,n−1] 中第一个大于 v=heights[a] 的高度的位置。这可以用线段树二分解决。

	创建一棵维护区间最大值 mx 的线段树。
	对于每个询问，递归这棵线段树，分类讨论：
	· 如果当前区间（线段树的节点对应的区间）最大值 mx≤v，则当前区间没有大于 v 的数，返回 −1。
	· 如果当前区间只包含一个元素，则找到答案，返回该元素的下标。
	· 如果左子树包含 b+1，则递归左子树。
	· 如果左子树返回 −1，则返回递归右子树的结果。

	注：方法三是最灵活的，如果题目还有动态修改 heights[i] 的操作，方法三也可以做。
*/

type seg []int

// 初始化线段树，维护去见最大值
func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 返回 [L,n-1] 中第一个 > v 的值的下标
// 如果不存在，返回 -1
func (t seg) query(o, l, r, L, v int) int {
	if t[o] <= v { // 区间最大值 <= v
		return -1 // 没有 > v 的数
	}
	if l == r { // 找到了
		return l
	}
	m := (l + r) >> 1
	if L <= m {
		pos := t.query(o<<1, l, m, L, v) // 递归左子树
		if pos >= 0 {                    // 找到了
			return pos
		}
	}
	return t.query(o<<1|1, m+1, r, L, v) // 递归右子树
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(heights)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(heights, 1, 0, n-1)

	ans := make([]int, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		if a > b {
			a, b = b, a // 保证 a <= b
		}
		if a == b || heights[a] < heights[b] {
			ans[i] = b // a 直接跳到 b
		} else {
			// 线段树二分，找 [b+1,n-1] 中第一个 > height[a] 的位置
			ans[i] = t.query(1, 0, n-1, b+1, heights[a])
		}
	}
	return ans
}
