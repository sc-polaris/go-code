package main

// 定长
func findAnagrams(s string, p string) (ans []int) {
	var cnt1 [26]int
	var cnt2 [26]int
	for _, c := range p {
		cnt2[c-'a']++
	}
	for r, c := range s {
		cnt1[c-'a']++
		l := r - len(p) + 1
		if l < 0 {
			continue
		}
		if cnt1 == cnt2 {
			ans = append(ans, l)
		}
		cnt1[s[l]-'a']--
	}
	return
}

// 不定长
func findAnagrams2(s string, p string) (ans []int) {
	var cnt [26]int
	for _, c := range p {
		cnt[c-'a']++
	}
	l := 0
	for r, c := range s {
		c -= 'a'
		cnt[c]--         // 右端点进
		for cnt[c] < 0 { // 字母 c 多了
			cnt[s[l]-'a']++ // 左端点字母离开窗口
			l++
		}
		if r-l+1 == len(p) {
			ans = append(ans, l)
		}
	}
	return
}
