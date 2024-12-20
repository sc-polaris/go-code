package main

/*
	给你一个字符串 s ，它由某个字符串 t 和若干 t  的 同位字符串 连接而成。

	请你返回字符串 t 的 最小 可能长度。

	同位字符串 指的是重新排列一个单词得到的另外一个字符串，原来字符串中的每个字符在新字符串中都恰好只使用一次。
*/

/*
	题意可知，字符串 t 的长度一定为字符串 s 的长度 n 的因数，因此我们可以从小到大枚举 n 的因数作为 t 的长度。令当前枚举的因数为 i，
	我们将字符串 s 切分为若干个长度为 i 的子字符串，用 count0 统计前一子字符串的字符出现次数，用 count1 统计后一子字符串的出现次
	数，如果 count0 不等于 count1，那么说明 i 不符合题意；否则说明所有子字符串的字符出现次数都相等，那么返回 i 作为 t 的最小可能长度。
*/

func minAnagramLength(s string) int {
	n := len(s)
	check := func(m int) bool {
		var count0 [26]int
		for j := 0; j < n; j += m {
			var count1 [26]int
			for k := j; k < j+m; k++ {
				count1[s[k]-'a']++
			}
			if j > 0 && count0 != count1 {
				return false
			}
			count0 = count1
		}
		return true
	}
	for i := 1; i < n; i++ {
		if n%i != 0 {
			continue
		}
		if check(i) {
			return i
		}
	}
	return n
}
