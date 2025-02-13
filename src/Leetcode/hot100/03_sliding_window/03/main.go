package main

func lengthOfLongestSubstring(s string) (ans int) {
	var cnt [128]int
	l := 0
	for r, c := range s {
		cnt[c]++
		for cnt[c] > 1 { // 窗口内有重复
			cnt[s[l]]--
			l++ // 缩小窗口
		}
		ans = max(ans, r-l+1)
	}
	return
}
