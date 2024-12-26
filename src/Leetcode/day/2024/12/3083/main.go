package main

/*
	给你一个字符串 s ，请你判断字符串 s 是否存在一个长度为 2 的子字符串，在其反转后的字符串中也出现。

	如果存在这样的子字符串，返回 true；如果不存在，返回 false 。
*/

func isSubstringPresent(s string) bool {
	vis := [26][26]bool{}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'a', s[i]-'a'
		vis[x][y] = true
		if vis[y][x] {
			return true
		}
	}
	return false
}

// 位运算优化

func isSubstringPresent2(s string) bool {
	vis := [26]int{}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'a', s[i]-'a'
		vis[x] |= 1 << y
		if vis[y]>>x&1 == 1 {
			return true
		}
	}
	return false
}
