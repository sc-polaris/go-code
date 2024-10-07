package main

import (
	"container/heap"
	"sort"
)

/*
	汽车从起点出发驶向目的地，该目的地位于出发位置东面 target 英里处。

	沿途有加油站，用数组 stations 表示。其中 stations[i] = [positioni, fueli] 表示第 i 个加油站位于出发位置东面 positioni 英里处，并且有 fueli 升汽油。

	假设汽车油箱的容量是无限的，其中最初有 startFuel 升燃料。它每行驶 1 英里就会用掉 1 升汽油。当汽车到达加油站时，它可能停下来加油，将所有汽油从加油站转移到汽车中。

	为了到达目的地，汽车所必要的最低加油次数是多少？如果无法到达目的地，则返回 -1 。

	注意：如果汽车到达加油站时剩余燃料为 0，它仍然可以在那里加油。如果汽车到达目的地时剩余燃料为 0，仍然认为它已经到达目的地。
*/

/*
	最大堆贪心
*/

func minRefuelStops(target int, startFuel int, stations [][]int) (ans int) {
	stations = append(stations, []int{target, 0})
	prePosition, curFuel := 0, startFuel
	fuelHeap := &hp{}
	for _, station := range stations {
		position, fuel := station[0], station[1]
		curFuel -= position - prePosition       // 每行驶 1 英里用掉 1 升汽油
		for fuelHeap.Len() > 0 && curFuel < 0 { // 没油了
			curFuel += heap.Pop(fuelHeap).(int) // 选油量最多的油桶
			ans++
		}
		if curFuel < 0 { // 无法到达
			return -1
		}
		heap.Push(fuelHeap, fuel) // 留着后面加油
		prePosition = position
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)         { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any           { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
