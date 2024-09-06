package main

import "strings"

func wordPattern(pattern string, s string) bool {
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i := range pattern {
		x, y := pattern[i], words[i]
		if yy, ok := m1[x]; ok && yy != y {
			return false
		}
		if xx, ok := m2[y]; ok && xx != x {
			return false
		}
		m1[x] = y
		m2[y] = x
	}
	return true
}
