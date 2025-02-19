package main

func wordBreak(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i == 0 {
			return 1
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if words[s[j:i]] && dfs(j) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(n) == 1
}

func wordBreak2(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if words[s[j:i]] && f[j] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}
