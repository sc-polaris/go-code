package binary_index_tree

import (
	"sort"
)

// 树状数组

type BinaryIndexTree struct {
	n int
	c []int
}

func NewBinaryIndexTree(n int) *BinaryIndexTree {
	c := make([]int, n+1)
	for i := range c {
		c[i] = -1
	}
	return &BinaryIndexTree{n: n, c: c}
}

func (t *BinaryIndexTree) update(x, v int) {
	for ; x <= t.n; x += x & -x {
		t.c[x] = max(t.c[x], v)
	}
}

func (t *BinaryIndexTree) query(x int) int {
	ans := -1
	for ; x > 0; x -= x & -x {
		ans = max(ans, t.c[x])
	}
	return ans
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n, m := len(nums1), len(queries)
	nums := make([][2]int, n)
	for i := range nums {
		nums[i] = [2]int{nums1[i], nums2[i]}
	}

	// 按照 nums1从大到小的顺序排序
	sort.Slice(nums, func(i, j int) bool { return nums[i][0] > nums[j][0] })
	sort.Ints(nums2)
	// queries[i]=(x,y)按照x从大到小排序
	idx := make([]int, m)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { return queries[idx[i]][0] > queries[idx[j]][0] })

	tree := NewBinaryIndexTree(n)
	ans := make([]int, m)
	j := 0
	for _, i := range idx {
		x, y := queries[i][0], queries[i][1]
		// 将所有大于等于x的元素的nums2的值插入树状数组中，
		// 树状数组维护的是离散化后的nums2的却见中nums1+nums2的最大值
		// 于树状数组维护的是前缀最大值，所以我们在实现上，可以将 nums2 反序插入到树状数组中。
		for ; j < n && nums[j][0] >= x; j++ {
			tree.update(n-sort.SearchInts(nums2, nums[j][1]), nums[j][0]+nums[j][1])
		}

		// 只需要在树状数组中查询大于等于离散化后的 y 区间对应的最大值即可
		ans[i] = tree.query(n - sort.SearchInts(nums2, y))
	}

	return ans
}
