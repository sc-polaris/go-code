package main

import (
	"container/heap"
	"sort"
)

/*
	给你一个整数数组 nums ，一个整数 k  和一个整数 multiplier 。

	你需要对 nums 执行 k 次操作，每次操作中：
	· 找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。
	· 将 x 替换为 x * multiplier 。
	k 次操作以后，你需要将 nums 中每一个数值对 109 + 7 取余。
	请你返回执行完 k 次乘运算之后，最终的 nums 数组。
*/

/*
	核心观察：对于两个数 x 和 y，如果 x 在 y 左边，且 x≤y 以及 x⋅multiplier>y，那么操作 y 之后，根据 x≤y，
	我们有 x⋅multiplier≤y⋅multiplier，这意味着下一次一定会操作 x。继续推导下去，后面的操作顺序是 y,x,y,x,⋯
	这意味着当两个数接近时，我们会交替操作这两个数，而不会连续操作同一个数。
	对于更多的数的情况也同理，当这些数接近时，我们会按照从小到大的顺序依次操作这些数。
	那么，首先用最小堆手动模拟操作，直到原数组的最大值 mx 成为这 n 个数的最小值。根据上面的结论，后面的操作就不需要手动模拟了。
	设此时还剩下 k 次操作，那么：
	· 对于前 k%n 小的数，还可以操作 ⌊k/n⌋+1次
	· 其余元素，还可以操作 ⌊k/n⌋ 次
	用快速幂计算操作这么多次后的结果，

*/

const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	heap.Init(&h)

	// 模拟，直到堆顶是 mx
	for ; k > 0 && h[0].x < mx; k-- {
		h[0].x *= multiplier
		heap.Fix(&h, 0)
	}

	// 剩余的操作可以直接用公式计算
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		e := k / n
		if i < k%n {
			e++
		}
		nums[p.i] = p.x % mod * pow(multiplier, e) % mod
	}
	return nums
}

type pair struct{ x, i int }

func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (_ any)      { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = (res * x) % mod
		}
		x = (x * x) % mod
	}
	return res
}
