package main

func longestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for j := range s0 {
		for _, s := range strs {
			if j == len(s) || s[j] != s0[j] {
				return s0[:j]
			}
		}
	}
	return s0
}
