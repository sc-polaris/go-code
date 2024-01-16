package main

import "strings"

func count1(num1 string, num2 string, minSum int, maxSum int) int {
	const mod = 1e9 + 7
	calc := func(s string) int {
		memo := make([][]int, len(s))
		for i := range memo {
			memo[i] = make([]int, min(9*len(s), maxSum)+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(int, int, bool) int
		dfs = func(i int, sum int, isLimit bool) (res int) {
			if sum > maxSum { // 非法
				return
			}
			if i == len(s) {
				if sum >= minSum { // 合法
					return 1
				}
				return
			}
			if !isLimit {
				p := &memo[i][sum]
				if *p >= 0 {
					return *p
				}
				defer func() { *p = res }()
			}
			up := 9
			if isLimit {
				up = int(s[i] - '0')
			}
			for d := 0; d <= up; d++ { // 枚举当前数位填 d
				res = (res + dfs(i+1, sum+d, isLimit && d == up)) % mod
			}
			return
		}
		return dfs(0, 0, true)
	}

	ans := calc(num2) - calc(num1) + mod // 避免负数
	sum := 0
	for _, c := range num1 {
		sum += int(c - '0')
	}
	if minSum <= sum && sum <= maxSum { // num1 是合法的，补回来
		ans++
	}
	return ans % mod
}

func count(num1 string, num2 string, minSum int, maxSum int) int {
	const mod = 1e9 + 7
	n := len(num2)
	num1 = strings.Repeat("0", n-len(num1)) + num1 // 补前导0，和 num2 对齐

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, min(9*n, maxSum)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool, bool) int
	dfs = func(i int, sum int, limitLow bool, limitHigh bool) (res int) {
		if sum > maxSum { // 非法
			return
		}
		if i == n {
			if sum >= minSum { // 合法
				return 1
			}
			return
		}

		if !limitLow && !limitHigh {
			p := &memo[i][sum]
			if *p > 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow {
			lo = int(num1[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(num2[i] - '0')
		}

		for d := lo; d <= hi; d++ { // 枚举当前位数填 d
			res = (res + dfs(i+1, sum+d, limitLow && d == lo, limitHigh && d == hi)) % mod
		}
		return
	}
	return dfs(0, 0, true, true)
}

func main() {

}
