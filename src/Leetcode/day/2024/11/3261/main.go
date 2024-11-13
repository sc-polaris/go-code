package main

import "sort"

/*
	给你一个 二进制 字符串 s 和一个整数 k。
	另给你一个二维整数数组 queries ，其中 queries[i] = [li, ri] 。
	如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：
	· 字符串中 0 的数量最多为 k。
	· 字符串中 1 的数量最多为 k。
	返回一个整数数组 answer ，其中 answer[i] 表示 s[li..ri] 中满足 k 约束 的子字符串的数量。
*/

/*
	方法一：滑动窗口+前缀和+二分查找
	核心思路：对于每个询问，计算以 l 为右端点的合法子串个数，以 l+1 为右端点的合法子串个数，……，以 r 为右端点的合法子串个数。

	我们需要知道以 i 为右端点的合法子串，其左端点最小是多少。

	由于随着 i 的变大，窗口内的字符数量变多，越不能满足题目要求，所以最小左端点会随着 i 的增大而增大，有单调性，因此可以用 滑动窗口 计算。

	设以 i 为右端点的合法子串，其左端点最小是 left[i]。
	那么以 i 为右端点的合法子串，其左端点可以是 left[i],left[i]+1,…,i，一共
						i−left[i]+1
	个。

	回答询问时，分类讨论：

	· 如果 left[r]≤l，说明 [l,r] 内的所有子串都是合法的，这一共有 1+2+⋯+(r−l+1)=(r−l+2)(r−l+1)/2 个。
	· 否则，由于 left 是有序数组，我们可以在 [l,r] 中 二分查找 left 中的第一个满足 left[j]≥l 的下标 j，那么：
		· 由于 left[j−1]<l，所以 [l,j−1] 内的所有子串都是合法的，这一共有 1+2+⋯+(j−l)=(j−l+1)(j−l)/2 个。
		· 右端点在 [j,r] 内的子串，可以累加下标在 [j,r] 内的所有 i−left[i]+1 的和。这可以用 前缀和 预处理。
		· 上述两项累加，即为答案。
	代码实现时，两种情况可以合并为一种。
*/

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l // 记录合法子串右端点 i 对应的最小左端点 l
		// 计算 i-left[i]+1 的前缀和
		sum[i+1] = sum[i] + i - l + 1
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := l + sort.SearchInts(left[l:r+1], l) // 如果区间内所有数都小于 l，结果是 j=r+1
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}

/*
	方法二：预处理
	上面的做法，每次都要二分找最小 j，满足 left[j]≥l。能否不用二分呢？

	也可以直接预处理，对每个左端点 l=0,1,2,…,n−1，计算出最小的 j，满足 left[j]≥l。

	由于 left 数组是有序的，这个过程可以用双指针实现。

	将计算出的 j 保存到 right[l] 中。如果不存在满足 left[j]≥l 的 j，则 right[l]=n。

	现在 left[right[l]]≥l 且 left[right[l]−1]<l。

	预处理后，回答询问时，j 可以直接通过 right[l] 获取到。注意这个数不能超过 r+1，所以有
						j=min(right[l],r+1)
	写法一
*/

func countKConstraintSubstrings2(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l // 记录合法子串右端点 i 对应的最小左端点 l
		// 计算 i-left[i]+1 的前缀和
		sum[i+1] = sum[i] + i - l + 1
	}

	right := make([]int, n)
	l = 0
	for i := range right {
		for l < n && left[l] < i {
			l++
		}
		right[i] = l
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := min(right[l], r+1)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}

/*
	写法二
	在上面的方法中，我们找的是最小的 j，满足 left[j]≥l。

	也可以找最小的 j，满足 left[j]>l，这不会影响计算出的合法子串个数。

	现在 left[right[l]]>l 且 left[right[l]−1]≤l。

	在滑窗的过程中，如果发现窗口不满足要求，那么在移动左端点 l 之前，可以记录 right[l]=i。
*/

func countKConstraintSubstrings3(s string, k int, queries [][]int) []int64 {
	n := len(s)
	right := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			right[l] = i
			l++
		}
		sum[i+1] = sum[i] + i - l + 1
	}
	// 剩余没填的 right[l] 均为 n
	for ; l < n; l++ {
		right[l] = n
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := min(right[l], r+1)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}
