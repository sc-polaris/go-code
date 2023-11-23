package sort

import (
	"sort"
)

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n, m := len(nums1), len(queries)
	a1, a2, a3 := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a1[i], a2[i], a3[i] = i, i, i
	}

	// nums1从大到小排序的索引
	sort.Slice(a1, func(i, j int) bool { return nums1[a1[i]] > nums1[a1[j]] })
	// nums2从大到小排序的索引
	sort.Slice(a2, func(i, j int) bool { return nums2[a2[i]] > nums2[a2[j]] })
	// nums1[i]+nums2[i]从大到小排序的索引
	sort.Slice(a3, func(i, j int) bool { return nums1[a3[i]]+nums2[a3[i]] > nums1[a3[j]]+nums2[a3[j]] })

	res := make([]int, m)
	for i := 0; i < m; i++ {
		maxV := -1
		x, y := queries[i][0], queries[i][1]
		for j := 0; j < n; j++ {
			v1, v2, v3 := a1[j], a2[j], a3[j]

			if x > nums1[v1] || y > nums2[v2] {
				break
			}

			if nums1[v1] >= x && nums2[v1] >= y && nums1[v1]+nums2[v1] > maxV {
				maxV = nums1[v1] + nums2[v1]
			}
			if nums1[v2] >= x && nums2[v2] >= y && nums1[v2]+nums2[v2] > maxV {
				maxV = nums1[v2] + nums2[v2]
			}
			if nums1[v3] >= x && nums2[v3] >= y {
				maxV = nums1[v3] + nums2[v3]
				break
			}
		}
		res[i] = maxV
	}

	return res
}
