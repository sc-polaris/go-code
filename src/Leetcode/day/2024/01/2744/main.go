package main

func maximumNumberOfStringPairs(words []string) (ans int) {
	st := [26][26]bool{}
	for _, s := range words {
		x, y := s[0]-'a', s[1]-'a'
		if st[y][x] {
			ans++ // s 和 st 中的 y+x 匹配
		} else {
			st[x][y] = true
		}
	}
	return
}

func main() {

}
