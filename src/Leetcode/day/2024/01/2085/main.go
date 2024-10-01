package main

func countWords(words1 []string, words2 []string) int {
	cnt1, cnt2 := make(map[string]int), make(map[string]int)
	for _, w := range words1 {
		cnt1[w]++
	}
	for _, w := range words2 {
		cnt2[w]++
	}

	res := 0
	for w, v := range cnt1 {
		if v == 1 && cnt2[w] == 1 {
			res++
		}
	}

	return res
}

func countWords2(words1 []string, words2 []string) int {
	cnt := make(map[string]int)
	for _, word := range words1 {
		cnt[word]++
	}

	res := 0
	for _, word := range words2 {
		if val, ok := cnt[word]; ok {
			if val == 1 {
				cnt[word] = -1
				res++
			}
			if val == -1 {
				cnt[word] = 0
				res--
			}
		}
	}
	return res
}

func main() {

}
