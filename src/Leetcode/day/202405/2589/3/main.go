package main

import (
	"slices"
	"sort"
)

/*
	方法三：栈+二分超找
	由于每次都是从右往左新增时间点，如果把连续的时间点看成闭区间，那么从右到左新增时间点，会把若干右侧的区间合并成一个大区间，也就是从 end 倒着
	开始，先合并右边，再合并左边，因此可以用栈来优化。

	栈中维护闭区间的左右端点，已经从栈底到栈顶的区间长度和（类似前缀和）。

	由于一旦发现区间相交就立即合并，所以栈中保存的都是不相交的区间。

	合并前，先尝试在栈中二分查找包含左端点的 start 区间。由于栈中还保存了区间长度之和，所以可以快速得到 [start,end] 范围内的运行中的时间点个数。

	如果还需要新增时间点，那么就从右到左合并。
*/

func findMinimumTime(tasks [][]int) int {
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
	// 栈中保存闭区间左右端点，栈底到栈顶的区间长度的和
	type tuple struct{ l, r, s int }
	st := []tuple{{-2, -2, 0}} // 哨兵，保证不和任何区间相交
	for _, p := range tasks {
		start, end, d := p[0], p[1], p[2]
		i := sort.Search(len(st), func(i int) bool { return st[i].l >= start }) - 1
		d -= st[len(st)-1].s - st[i].s // 去掉运行中的时间点
		if start <= st[i].r {          // start 在区间 st[i] 内
			d -= st[i].r - start + 1 // 去掉运行中的时间点
		}
		if d <= 0 {
			continue
		}
		for end-st[len(st)-1].r <= d { // 剩余的 d 填充区间后缀
			top := st[len(st)-1]
			st = st[:len(st)-1]
			d += top.r - top.l + 1 // 合并区间
		}
		st = append(st, tuple{end - d + 1, end, st[len(st)-1].s + d})
	}
	return st[len(st)-1].s
}
