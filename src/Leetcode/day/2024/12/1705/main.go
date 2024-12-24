package _705

import "container/heap"

/*
	有一棵特殊的苹果树，一连 n 天，每天都可以长出若干个苹果。在第 i 天，树上会长出 apples[i] 个苹果，
	这些苹果将会在 days[i] 天后（也就是说，第 i + days[i] 天时）腐烂，变得无法食用。也可能有那么几天，
	树上不会长出新的苹果，此时用 apples[i] == 0 且 days[i] == 0 表示。

	你打算每天 最多 吃一个苹果来保证营养均衡。注意，你可以在这 n 天之后继续吃苹果。

	给你两个长度为 n 的整数数组 days 和 apples ，返回你可以吃掉的苹果的最大数目。
*/

/*
	核心思路：对于两个苹果 A 和 B，设 A 更早腐烂，那么应该先吃 A。如果先吃 B，可能下一天 A 就烂了。

	用最小堆维护苹果的腐烂日期和个数，模拟一天吃一个苹果。

	注意第 n 天之后还可以继续吃苹果，所以要一直模拟到堆为空为止。
*/

func eatenApples(apples []int, days []int) (ans int) {
	h := hp{}
	for i := 0; i < len(apples) || h.Len() > 0; i++ {
		for h.Len() > 0 && h[0].rottenDay == i { // 已经腐烂
			heap.Pop(&h)
		}
		if i < len(apples) && apples[i] > 0 {
			heap.Push(&h, pair{i + days[i], apples[i]})
		}
		if h.Len() > 0 {
			// 吃一个最早腐烂的苹果
			ans++
			h[0].num--
			if h[0].num == 0 {
				heap.Pop(&h)
			}
		}
	}
	return
}

/*
	当 i≥n 时，我们浪费太多时间在「一个一个地」吃苹果上了。事实上，可以直接把最早腐烂的苹果（在腐烂前）一次性解决掉，也就是直接把 i 跳到吃完（或者腐烂）的那一天。

	设当前是第 i 天，最早腐烂的苹果有 num 个，在第 rottenDay 天腐烂。

	那么我们可以在腐烂前吃掉

	k=min(num,rottenDay−i)
	个苹果。

	吃完后，把答案和 i 都增加 k，继续循环，直到堆为空为止。
*/

func eatenApples2(apples, days []int) (ans int) {
	h := hp{}
	for i, num := range apples {
		for h.Len() > 0 && h[0].rottenDay == i { // 已腐烂
			heap.Pop(&h)
		}
		if num > 0 {
			heap.Push(&h, pair{i + days[i], num})
		}
		if h.Len() > 0 {
			// 吃一个最早腐烂的苹果
			ans++
			h[0].num--
			if h[0].num == 0 {
				heap.Pop(&h)
			}
		}
	}

	i := len(apples)
	for {
		for h.Len() > 0 && h[0].rottenDay <= i { // 已腐烂
			heap.Pop(&h)
		}
		if h.Len() == 0 {
			return
		}
		p := heap.Pop(&h).(pair)
		k := min(p.num, p.rottenDay-i)
		ans += k
		i += k
	}
	return
}

type pair struct{ rottenDay, num int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].rottenDay < h[j].rottenDay }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
