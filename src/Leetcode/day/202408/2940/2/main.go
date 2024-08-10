package main

import "sort"

/*
	方法二：离线+单调栈二分

	同方法一，先遍历 queries，处理出 qs。

	然后倒序遍历 heights。试想一下，如果 heights[2]=8,heights[3]=6，那么对于在 heights[2] 左边的高度来说，heights[3] 必然不是第一个相遇的位置，
	因为我们总是可以选择比 heights[3] 更大且更靠左的 heights[2]。这意味着，当我们遍历到一个更大的高度时，之前遍历过的更小的高度就是无用数据了，要及时清除掉。

	这启发我们用一个底大顶小的单调栈维护高度。

	由于栈中高度严格递减（从栈底到栈顶），可以二分查找最后一个大于 heights[a] 的高度。

	代码实现时，为方便计算下标，栈中保存的是高度的下标。
*/

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	type pair struct{ h, i int }
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

	var st []int
	for i := len(heights) - 1; i >= 0; i-- {
		for _, q := range qs[i] {
			j := sort.Search(len(st), func(i int) bool { return heights[st[i]] <= q.h }) - 1
			if j >= 0 {
				ans[q.i] = st[j]
			} else {
				ans[q.i] = -1
			}
		}
		for len(st) > 0 && heights[i] >= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}

/*
private int binarySearch(int[] heights, int[] st, int right, int x) {
	int left = -1; // 开区间 (left, right)
	while (left + 1 < right) { // 开区间不为空
		int mid = (left + right) >>> 1;
		if (heights[st[mid]] > x) {
			left = mid; // 范围缩小到 (mid, right)
		} else {
			right = mid; // 范围缩小到 (left, mid)
		}
	}
	return left < 0 ? -1 : st[left];
}
*/
