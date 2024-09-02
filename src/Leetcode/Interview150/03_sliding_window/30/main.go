package main

/*
	给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
	s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
	· 例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"，
      和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。

	返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。
*/

func findSubstring(s string, words []string) (res []int) {
	sLen, wLen, k := len(s), len(words), len(words[0])
	if sLen < k*wLen {
		return []int{}
	}

	mp := make(map[string]int)
	for _, w := range words {
		mp[w]++
	}

	// 移动窗口减少重复检查单词，按照单词长度取不同批次
	for i := 0; i < k; i++ {
		count := 0
		cMap := make(map[string]int)
		for l, r := i, i; r <= sLen-k; r += k {
			word := s[r : r+k]
			if num, ok := mp[word]; ok {
				// 如果计数器中单词数目超标，左移指针直至符合数目要求
				for ; cMap[word] >= num; l += k {
					cMap[s[l:l+k]]--
					count--
				}
				cMap[word]++
				count++
			} else {
				// 如果当前单词不在词典里，左移指针至下一个单词，左移过程中清理计数
				for ; l < r; l += k {
					cMap[s[l:l+k]]--
					count--
				}
				l += k
			}
			if count == wLen {
				res = append(res, l)
			}
		}
	}

	return
}
