package main

func lengthOfLongestSubstring(s string) (ans int) {
	cnt := [128]bool{}
	l := 0
	for r, ch := range s {
		for cnt[ch] {
			cnt[s[l]] = false
			l++
		}
		cnt[ch] = true
		ans = max(ans, r-l+1)
	}
	return
}
