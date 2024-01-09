package main

/*
	f[i] 表示字符串 s 的前 i 个字符的最小额外字符数，f[0] = 0
*/

func minExtraChar(s string, dictionary []string) int {
	st := make(map[string]bool)
	for _, s := range dictionary {
		st[s] = true
	}
	n := len(s)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] + 1        // 不选
		for j := 0; j < i; j++ { // 枚举选哪个
			if st[s[j:i]] {
				f[i] = min(f[i], f[j])
			}
		}
	}

	return f[n]
}

func main() {

}
