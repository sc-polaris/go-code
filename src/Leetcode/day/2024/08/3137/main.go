package main

/*
	给你一个长度为 n 的字符串 word 和一个整数 k ，其中 k 是 n 的因数。

	在一次操作中，你可以选择任意两个下标 i 和 j，其中 0 <= i, j < n ，且这两个下标都可以被 k 整除，然后用从 j 开始的长度为 k 的子串替换从 i 开始的长度为 k 的子串。
	也就是说，将子串 word[i..i + k - 1] 替换为子串 word[j..j + k - 1] 。

	返回使 word 成为 K 周期字符串 所需的 最少 操作次数。

	如果存在某个长度为 k 的字符串 s，使得 word 可以表示为任意次数连接 s ，则称字符串 word 是 K 周期字符串 。例如，如果 word == "ababab"，那么 word 就是 s = "ab"
	时的 2 周期字符串 。
*/

/*
	根据题意，我们只能选择首字母下标为 0,k,2k,3k,⋯,n−k 的长为 k 的子串来操作（替换）。
	并且，k 周期字符串意味着，所有首字母下标为 0,k,2k,3k,⋯,n−k 的长为 k 的子串均相等。

	为使操作次数尽量少，我们可以计算最多保留多少个子串不变。也就是统计 word 中的这些首字母下标为 0,k,2k,3k,⋯,n−k 的长为 k 的子串中，出现次数最多的子串的出现次数 mx。
	用出现次数最多的子串，替换其余子串。

	所以用子串的个数 n/k 减去 mx，就是最少操作次数。
*/

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	cnt := make(map[string]int)
	for i := k; i <= n; i += k {
		cnt[word[i-k:i]]++
	}
	mx := 0
	for _, c := range cnt {
		mx = max(mx, c)
	}
	return n/k - mx
}
