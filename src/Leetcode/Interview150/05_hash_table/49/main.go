package main

import "slices"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		t := []byte(s)
		slices.Sort(t)
		sortedS := string(t)
		m[sortedS] = append(m[sortedS], s)
	}
	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
