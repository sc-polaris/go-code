package main

import "strings"

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStrKmp(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	// 求 needle 的前缀函数
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
