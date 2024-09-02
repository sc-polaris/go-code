package main

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	i := 0
	for _, c := range t {
		if s[i] == byte(c) {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}
