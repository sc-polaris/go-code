package main

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	var ans []byte
	k := 0
	for i, j := 25, 24; i >= 0 && j >= 0; {
		switch {
		case cnt[i] == 0: // 当前字符已经填完，填入后面的字符，重置 k
			k, i = 0, i-1
		case k < repeatLimit: // 当前字符未超过限制
			cnt[i]--
			ans = append(ans, 'a'+byte(i))
			k++
		case j >= i || cnt[j] == 0: // 当前字符已经超过限制，查找可填入的其他字符
			j--
		default: // 当前字符已经超过限制，填入其他字符，并且重置 k
			cnt[j]--
			ans = append(ans, 'a'+byte(j))
			k = 0
		}
	}

	return string(ans)
}

func repeatLimitedString2(s string, repeatLimit int) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	var ans []byte
	for i := 25; i >= 0; i-- { // 从大到小填字母
		k := i - 1
		for {
			// 填充 min(repeatLimit, cnt[i])
			for j := 0; j < repeatLimit && cnt[i] > 0; j++ {
				cnt[i]--
				ans = append(ans, 'a'+byte(i))
			}
			if cnt[i] == 0 { // i 用完了，找下一个更小的
				break
			}
			for k >= 0 && cnt[k] == 0 {
				k--
			}
			if k < 0 {
				break
			}
			// i 没用完，插入一个字母 k，继续填 i
			cnt[k]--
			ans = append(ans, 'a'+byte(k))
		}
	}

	return string(ans)
}

func main() {

}
