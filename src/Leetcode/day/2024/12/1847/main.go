package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math"
	"slices"
)

/*
	一个酒店里有 n 个房间，这些房间用二维整数数组 rooms 表示，其中 rooms[i] = [roomIdi, sizei] 表示有一个房间号为 roomIdi 的房间且它的面积为 sizei 。每一个房间号 roomIdi 保证是 独一无二 的。

	同时给你 k 个查询，用二维数组 queries 表示，其中 queries[j] = [preferredj, minSizej] 。第 j 个查询的答案是满足如下条件的房间 id ：
	房间的面积 至少 为 minSizej ，且
	· abs(id - preferredj) 的值 最小 ，其中 abs(x) 是 x 的绝对值。
	· 如果差的绝对值有 相等 的，选择 最小 的 id 。如果 没有满足条件的房间 ，答案为 -1 。

	请你返回长度为 k 的数组 answer ，其中 answer[j] 为第 j 个查询的结果。
*/

/*
	核心思路
	把询问排序，通过改变回答询问的顺序，使问题更容易处理。
	比如有两个询问，其中 minSize 分别为 3 和 6。
	我们可以先回答 minSize=6 的询问，再回答 minSize=3 的询问。
	也就是先把面积 ≥6 的房间号添加到一个有序集合中，回答 minSize=6 的询问；然后把面积 ≥3 的房间号添加到有序集合中，回答 minSize=3 的询问。
	这里的关键是，由于面积 ≥6 的房间编号已经添加到有序集合中了，所以后续只需把面积在 [3,5] 中的房间号添加到有序集合中，不需要重复处理面积 ≥6 的房间。

	具体思路
	直接对 queries 排序是不行的，因为返回的答案必须按照询问的顺序。

	解决办法：设 q 是 queries 的长度，创建一个下标数组 queryIds=[0,1,2,…,q−1]，把下标根据 queries 的 minSize 从大到小排序，这样就避免直接对 queries 排序了。

	把 rooms 按照 size 从小到大排序（也可以从大到小）。

	然后创建一个有序集合 roomIds。用双指针遍历 queryIds 和 rooms，把房间面积 ≥minSize 的房间号添加到 roomIds 中。
	然后在 roomIds 中搜索离 preferred 最近的左右两个房间号，其中离 preferred 最近的房间号就是答案
*/

func closestRoom(rooms [][]int, queries [][]int) []int {
	// 按照 size 从大到小排序
	slices.SortFunc(rooms, func(a, b []int) int { return b[1] - a[1] })

	q := len(queries)
	queryIds := make([]int, q)
	for i := range queryIds {
		queryIds[i] = i
	}
	// 按照 minSize 从大到小排序
	slices.SortFunc(queryIds, func(i, j int) int { return queries[j][1] - queries[i][1] })

	ans := make([]int, q)
	for i := range ans {
		ans[i] = -1
	}
	roomIds := redblacktree.New[int, struct{}]()
	j := 0
	for _, i := range queryIds {
		preferredId, minSize := queries[i][0], queries[i][1]
		for j < len(rooms) && rooms[j][1] >= minSize {
			roomIds.Put(rooms[j][0], struct{}{})
			j++
		}

		diff := math.MaxInt
		// 左边的差
		if node, ok := roomIds.Floor(preferredId); ok {
			diff = preferredId - node.Key
			ans[i] = node.Key
		}
		// 右边的差
		if node, ok := roomIds.Ceiling(preferredId); ok && node.Key-preferredId < diff {
			ans[i] = node.Key
		}
	}
	return ans
}
