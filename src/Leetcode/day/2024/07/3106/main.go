package main

/*
	给你一个字符串 s 和一个整数 k 。

	定义函数 distance(s1, s2) ，用于衡量两个长度为 n 的字符串 s1 和 s2 之间的距离，即：
	· 字符 'a' 到 'z' 按 循环 顺序排列，对于区间 [0, n - 1] 中的 i ，计算所有「 s1[i] 和 s2[i] 之间 最小距离」的 和 。
	例如，distance("ab", "cd") == 4 ，且 distance("a", "z") == 1 。

	你可以对字符串 s 执行 任意次 操作。在每次操作中，可以将 s 中的一个字母 改变 为 任意 其他小写英文字母。

	返回一个字符串，表示在执行一些操作后你可以得到的 字典序最小 的字符串 t ，且满足 distance(s, t) <= k 。
*/

/*
	题意解读：
	每次操作可以把 s[i] 加一或减一，求在操作次数不超过 k 的前提下，s 的最小字典序。注意 z 加一后变成 a，a 减一后变成 z。

	分析：
	贪心，优先把左边的字母变成 a。
	把 s[i] 变成 a，可以不断减一到 a，也可以不断加一到 a，二者取最小值，得操作次数
						dis = min(s[i]-a,z-s[i]+1)
	算法：
	1. 从左到右遍历 s。
	2. 如果把 s[i] 变成 a 的操作次数 dis ≤ k，那么就把 s[i] 变成 a，同时 k 减少 dis。
	3. 否则无法变成 a，直接把 s[i] 减少 k，退出循环。
*/

func getSmallestString(s string, k int) string {
	t := []byte(s)
	for i, c := range t {
		dis := int(min(c-'a', 'z'-c+1))
		if dis > k {
			t[i] -= byte(k)
			break
		}
		t[i] = 'a'
		k -= dis
	}
	return string(t)
}
