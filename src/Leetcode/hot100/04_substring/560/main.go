package main

/*
	s[j]−s[i]=k (i<j)
	s[i]=s[j]−k
	计算 s[i] 的个数，等价于计算在 s[j] 左边的 s[j]−k 的个数。
*/

func subarraySum(nums []int, k int) (ans int) {
	s := 0
	cnt := map[int]int{0: 1}
	for _, x := range nums {
		s += x
		ans += cnt[s-k]
		cnt[s]++
	}
	return
}
