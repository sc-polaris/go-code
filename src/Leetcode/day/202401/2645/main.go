package main

/*
	1. 考虑相邻字母
	设x=s[i-1],y=s[i]
	若s有效，插入 y-x-1个
	可能为负数 (y-x-1+3)%3 则 (y-x+2)%3
	开始：s[0] != 'a' s前面插入s[0]-'a'个字母 s[n-1]!='c'插入'c'-s[n-1]
    合并： s[0]-'a'+'c'-s[n-1]=s[0]-s[n-1]+2

	2. 考虑abc的个数
	设 答案由 t 个 "abc" 组成 需要插入 3t -n
    if x < y  x和y可以在同一个"abc"
    if x >= y 一定不在
*/

func addMinimum1(s string) int {
	ans := int(s[0] - s[len(s)-1] + 2)
	for i := 1; i < len(s); i++ {
		ans += int(s[i]-s[i-1]+2) % 3
	}
	return ans
}

func addMinimum(s string) int {
	t, n := 1, len(s)
	for i := 1; i < n; i++ {
		if s[i-1] >= s[i] { // 一定有一个新的 abc
			t++
		}
	}
	return 3*t - n
}

func main() {

}
