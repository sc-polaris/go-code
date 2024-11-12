package main

/*
	给你一个 二进制 字符串 s 和一个整数 k。
	如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：
	· 字符串中 0 的数量最多为 k。
	· 字符串中 1 的数量最多为 k。
	返回一个整数，表示 s 的所有满足 k 约束 的子字符串的数量。
*/

func countKConstraintSubstrings(s string, k int) (ans int) {
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		ans += i - l + 1
	}
	return
}
