package main

import "fmt"

func makeSmallestPalindrome(s string) string {
	t := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		t[l] = min(s[l], s[r])
		t[r] = t[l]
	}
	return string(t)
}

func main() {
	fmt.Println(makeSmallestPalindrome("egcfe"))
}
