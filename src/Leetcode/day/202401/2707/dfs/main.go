package main

/*
	1. 跳过 s 的最后一个字符，那么问题变成 s 的前 n-1 个字符的子问题
	2. 考虑「枚举选那个」，如果从 s[j] 开始的后缀在 dictionary 中，那么问题变成 s 的前 j-1 个字符的子问题

	dfs(i) 表示前 i 个字符采取最优策略分割后，剩余的最少字符
	1. 跳过最后一个字符，有 dfs(i)=dfs(i-1)+1
	2. 考虑「枚举选哪个」，如果从 s[j] 到 s[i] 的子串在 dictionary 中，有 dfs(i)=min dfs(j) (j=0,1,...i-1)
	递归边界 dfs(0)=0
*/

func minExtraChar(s string, dictionary []string) int {
	st := make(map[string]bool)
	for _, s := range dictionary {
		st[s] = true
	}
	n := len(s)
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}

		// 不选
		res := dfs(i-1) + 1
		// 枚举选哪个
		for j := 0; j < i; j++ {
			if st[s[j:i]] {
				res = min(res, dfs(j))
			}
		}

		*p = res
		return res
	}
	return dfs(n)
}

func main() {
}
