package main

import (
	"maps"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		t := []byte(s)
		slices.Sort(t)
		sortedS := string(t)
		m[sortedS] = append(m[sortedS], s) // sortedS 相同的字符串分到同一组
	}
	return slices.Collect(maps.Values(m))
}
