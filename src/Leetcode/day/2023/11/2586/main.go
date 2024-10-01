package main

import "slices"

func vowelStrings(words []string, left int, right int) (res int) {
	chs := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, word := range words[left : right+1] {
		if slices.Contains(chs, word[0]) && slices.Contains(chs, word[len(word)-1]) {
			res++
		}
	}
	return
}

func main() {

}
