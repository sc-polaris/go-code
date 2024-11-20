package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	type pair struct {
		word string
		step int
	}
	q := []pair{{beginWord, 1}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		word := p.word
		for i := range word {
			for j := 'a'; j <= 'z'; j++ {
				nxt := word[:i] + string(j) + word[i+1:]
				if wordSet[nxt] {
					if nxt == endWord {
						return p.step + 1
					}
					wordSet[nxt] = false
					q = append(q, pair{nxt, p.step + 1})
				}
			}
		}
	}
	return 0
}
