package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	st, ed := 0, 0
	for i := range s {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r1-l1 > ed-st {
			st, ed = l1, r1
		}
		if r2-l2 > ed-st {
			st, ed = l2, r2
		}
	}
	return s[st : ed+1]
}

func expandAroundCenter(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}
