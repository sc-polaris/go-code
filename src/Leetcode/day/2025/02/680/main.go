package main

/*
	给你一个字符串 s，最多 可以从中删除一个字符。

	请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。
*/

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			// 删除 s[i] 或者 s[j]（注意 Go 的切片是 O(1) 的，不会生成新字符串）
			return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j])
		}
		i++
		j--
	}
	return true
}
