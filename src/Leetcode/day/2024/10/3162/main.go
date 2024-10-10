package main

/*
	给你两个整数数组 nums1 和 nums2，长度分别为 n 和 m。同时给你一个正整数 k。

	如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对（0 <= i <= n - 1, 0 <= j <= m - 1）。

	返回 优质数对 的总数。
*/

// 枚举因子
func numberOfPairs(nums1 []int, nums2 []int, k int) (ans int) {
	cnt := make(map[int]int)
	for _, x := range nums1 {
		if x%k > 0 {
			continue
		}
		x /= k
		for d := 1; d*d <= x; d++ { // 枚举因子
			if x%d == 0 {
				cnt[d]++ // 统计因子
				if d*d < x {
					cnt[x/d]++ // 因子总是成对出现
				}
			}
		}
	}
	for _, x := range nums2 {
		ans += cnt[x]
	}
	return
}

/*
	枚举倍数
	例如 b[j]=3，枚举 3,6,9,12,⋯，统计 a 中有多少个 a[i]/k 等于 3,6,9,12,⋯

	1. 统计 a[i]/k 的出现次数，保存到哈希表 cnt1 中。
	2. 统计 b[j] 的出现次数（相同 b[j] 无需重复计算），保存到哈希表 cnt2 中
	3. 设 cnt1 的最大key 为 u。
	4. 枚举 cnt2 中的元素 x，然后枚举 x 的倍数 y=x,2x,3x,⋯（不超过 u），累加 cnt1[y]，再乘上 cnt2[x]，加入答案。
*/

func numberOfPairs2(nums1 []int, nums2 []int, k int) (ans int) {
	cnt1 := make(map[int]int)
	u := 0
	for _, x := range nums1 {
		if x%k == 0 {
			u = max(u, x/k)
			cnt1[x/k]++
		}
	}
	if u == 0 {
		return
	}

	cnt2 := make(map[int]int)
	for _, x := range nums2 {
		cnt2[x]++
	}

	for x, cnt := range cnt2 {
		s := 0
		for y := x; y <= u; y += x { // 枚举 x 的倍数
			s += cnt1[y]
		}
		ans += s * cnt
	}
	return
}
