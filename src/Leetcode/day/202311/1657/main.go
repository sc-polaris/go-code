package main

import (
	"slices"
)

func closeStrings(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sMask, tMask int
	var sCnt, tCnt [26]int
	for _, c := range s {
		sMask |= 1 << (c - 'a') // 记录 s 总有字符c
		sCnt[c-'a']++
	}
	for _, c := range t {
		tMask |= 1 << (c - 'a') // 记录 t 总有字符c
		tCnt[c-'a']++
	}

	slices.Sort(sCnt[:])
	slices.Sort(tCnt[:])
	return sMask == tMask && slices.Equal(sCnt[:], tCnt[:])
}

func main() {

}
