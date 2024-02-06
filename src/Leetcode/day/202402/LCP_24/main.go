package main

import (
	"container/heap"
	"sort"
)

/*
	1. 转化
		设 nums 的前 5 个数分别是 a0,a1,a2,a3,a4。假设 a0 最终变成整数 x，那么这 5 个数 最终要变成
		x,x+1,x+2,x+3,x+4，作次数为
				|a0-x|+|a1-(x+1)|+|a2-(x+2)|+|a3-(x+3)|+|a4-(x+4)|
			=	[a0-x|+|(a1-1)-x|+|(a2-2)-x|+|(a3-3)-x|+|(a4-4)-x|
		最小化上面的值：
		定义 bi = ai-i, 问题变成求
				|b0-x|+|b1-x|+|b2-x|+|b3-x|+|b4-x|
		的最小值

	2. 中位数贪心
		把 bi 画在数轴上，|bi-x|叫做「bi 到 x 的距离」，现在求所有 bi 到 x 的最小距离和
		为方便计算，把 b 从小到大排序
		定理：将 b 的所有元素变成 b 的中位数是最优的。
		分类讨论：
			如果 n 是偶数，那么最小距离和等于 b 的右半之和减去左半之和，即
				(bn-1 - b0)+(bn-2 -b1)+...+(bn/2 - bn/2-1)
			=	(bn/2+bn/2+1+...+bn-1)-(b0+b1+...+bn/2-1)
			如果 n 是奇数，最小距离和等于去掉中位数 bn/2 向下取整之后，b 的右半之和减去左半之和

	3. 对顶堆维护动态中位数
		用一个大根堆 left 维护较小的一半，其元素和为 leftSum；用一个小根堆 right 维护较大的一般，其元素和
		是 rightSum，便利 nums[i]，设 b = nums[i]-i，分类讨论：
		1. 如果前缀长度是奇数，此时 left 和 right 大小相等，我们先把 b 插入 left，然后弹出 left 的堆顶，
			加到 right 中。这一操作可以保证，无论 b 是大是小，此时 right 的堆顶就是中位数 x。
			最小距离和为 rightSum - x - leftSum
		2. 如果前缀长度是偶数，此时 left 就比 right 少一个元素，我们先把 b 插入 right，然后弹出 right 的
			堆顶，加到 left 中。最小距离和为 rightSum - leftSum
		最后，对 1e9+7 取 mod
*/

type hp struct {
	sort.IntSlice     // 继承 Len，Less，Swap
	sum           int // 堆中元素之合
}

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)); h.sum += v.(int) }
func (*hp) Pop() (_ any) { return } // 没用到，无需实现

// pushPop 先把 v 入堆，然后弹出并返回堆顶
// 如果 v <= 堆顶，则直接返回 v
func (h *hp) pushPop(v int) int {
	if h.Len() > 0 && v > h.IntSlice[0] {
		h.sum += v - h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}

func numsGame(nums []int) []int {
	const mod = 1e9 + 7
	ans := make([]int, len(nums))
	left := &hp{}  // 维护较小的一半，大根堆（小根堆取负号）
	right := &hp{} // 维护较大的一半，小根堆
	for i, b := range nums {
		b -= i
		if i&1 == 0 {
			heap.Push(right, -left.pushPop(-b))
			x := right.IntSlice[0] // 中位数
			// 原本要减去 left.sum，但由于 left 所有元素都取负号
			ans[i] = (right.sum - x + left.sum) % mod
		} else {
			heap.Push(left, -right.pushPop(b))
			ans[i] = (right.sum + left.sum) % mod
		}
	}
	return ans
}

func main() {

}
