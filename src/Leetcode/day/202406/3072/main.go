package main

import (
	"slices"
	"sort"
)

/*
	给你一个下标从 1 开始、长度为 n 的整数数组 nums 。

	现定义函数 greaterCount ，使得 greaterCount(arr, val) 返回数组 arr 中 严格大于 val 的元素数量。

	你需要使用 n 次操作，将 nums 的所有元素分配到两个数组 arr1 和 arr2 中。在第一次操作中，将 nums[1] 追加到 arr1 。在第二次操作中，将 nums[2] 追加到 arr2 。之后，在第 i 次操作中：

	· 如果 greaterCount(arr1, nums[i]) > greaterCount(arr2, nums[i]) ，将 nums[i] 追加到 arr1 。
	· 如果 greaterCount(arr1, nums[i]) < greaterCount(arr2, nums[i]) ，将 nums[i] 追加到 arr2 。
	· 如果 greaterCount(arr1, nums[i]) == greaterCount(arr2, nums[i]) ，将 nums[i] 追加到元素数量较少的数组中。
	· 如果仍然相等，那么将 nums[i] 追加到 arr1 。
	连接数组 arr1 和 arr2 形成数组 result 。例如，如果 arr1 == [1,2,3] 且 arr2 == [4,5,6] ，那么 result = [1,2,3,4,5,6] 。

	返回整数数组 result 。
*/

type fenwick []int

// 把下标为 i 的元素增加 v
func (f fenwick) add(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] += v
	}
}

// 返回下标在 [1,i] 的元素之和
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func resultArray(nums []int) (ans []int) {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)

	a := nums[:1]
	b := []int{nums[1]}
	t := make(fenwick, m+1)
	t.add(m-sort.SearchInts(sorted, nums[0]), 1)
	t.add(m-sort.SearchInts(sorted, nums[1]), -1)
	for _, x := range nums[2:] {
		v := m - sort.SearchInts(sorted, x)
		d := t.pre(v - 1) // 转换成 < v 的元素个数之差
		if d > 0 || d == 0 && len(a) <= len(b) {
			a = append(a, x)
			t.add(v, 1)
		} else {
			b = append(b, x)
			t.add(v, -1)
		}
	}
	return append(a, b...)
}
