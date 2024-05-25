package main

/*
	给你一个整数数组 nums 和一个正整数 k ，返回长度为 k 且最具 竞争力 的 nums 子序列。

	数组的子序列是从数组中删除一些元素（可能不删除元素）得到的序列。

	在子序列 a 和子序列 b 第一个不相同的位置上，如果 a 中的数字小于 b 中对应的数字，那么
	我们称子序列 a 比子序列 b（相同长度下）更具 竞争力 。 例如，[1,3,4] 比 [1,3,5] 更具
	竞争力，在第一个不相同的位置，也就是最后一个位置上， 4 小于 5 。
*/

func mostCompetitive(nums []int, k int) []int {
	st := nums[:0] // 把 nums 当作栈
	for i, x := range nums {
		for len(st) > 0 && x < st[len(st)-1] && len(st)+len(nums)-i > k {
			st = st[:len(st)-1]
		}
		if len(st) < k {
			st = append(st, x)
		}
	}
	return st
}
