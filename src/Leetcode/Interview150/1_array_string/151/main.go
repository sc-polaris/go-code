package main

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	res := strings.Fields(s)
	slices.Reverse(res)
	return strings.Join(res, " ")
}
