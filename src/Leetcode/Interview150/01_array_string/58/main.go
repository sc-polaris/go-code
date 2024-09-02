package main

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	return len(s) - 1 - strings.LastIndex(s, " ")
}
