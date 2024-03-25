package main

/*
	dfs(i,c) 表示用前 i 种硬币组成金额 c 的方案数，选 or 不选
*/

func change(amount int, coins []int) int {
	n := len(coins)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return
		}
		p := &memo[i][c]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i, c-coins[i])
	}
	return dfs(n-1, amount)
}

func main() {

}
