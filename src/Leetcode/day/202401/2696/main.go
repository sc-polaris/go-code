package main

import "strings"

func minLength1(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}
	return len(s)
}

func minLength(s string) int {
	var st []byte
	for _, c := range s {
		if len(st) > 0 && (c == 'B' && st[len(st)-1] == 'A' || c == 'D' && st[len(st)-1] == 'C') {
			st = st[:len(st)-1]
		} else {
			st = append(st, byte(c))
		}
	}
	return len(st)
}

func main() {

}
