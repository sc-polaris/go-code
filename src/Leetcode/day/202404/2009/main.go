package main

import (
	"slices"
)

/*
	正难则反，考虑最多保留多少个元素不变。

	设 x 是修改后的连续数字的最大值，则修改后的连续数字的范围为闭区间 [x-n+1, x]，其中 n
	是 nums 的长度。在修改前，对于已经在 [x-n+1, x] 中的数，我们无需修改。那么，x 取多少，
	可以让无需修改的数最多呢？

	由于元素的位置不影响答案，且要求所有元素互不相同，我们可以将 nums 从小到大排序，并去掉
	重复元素。设 a 为 nums 排序去重后的数组。把 a[i] 画在一条数轴上，本题相当于又一个长度
	为 n 的滑动窗口，我们需要计算窗口内最多可以包含多少个数轴上的点。

	定理：只需要枚举 a[i] 作为窗口的右端点。
	证明：在窗口从左向右滑动的过程中，如果窗口右端点处没有点，那么继续滑动，在滑到下一个点之
         前，窗口内包含的点的个数是不会增多的。

	为了算出窗口内有多少个点，我们需要知道窗口包含的最左边的点在哪，设这个点的位置是 a[left]，
	则它必须大于等于窗口的左边界，即
								a[left] >= a[i]-n+1
	此时窗口内右 i-left+1 个点，取最大值，就得到了最多保留不变的元素个数。最后用 n 减去保留
	不变的元素个数，就得到了答案。
*/

func minOperations(nums []int) int {
	n := len(nums)
	slices.Sort(nums)
	a := slices.Compact(nums)
	ans, left := 0, 0
	for i, x := range a {
		for a[left] < x-n+1 { // a[left] 不在窗口内
			left++
		}
		ans = max(ans, i-left+1)
	}
	return n - ans
}
