package main

// 定义dfs(i,j)表示能否从nums[0]到nums[i]中选出一个和恰好等于 j的子序列。
func canPartition(nums []int) bool {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 != 0 {
		return false
	}
	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, s/2+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i < 0 {
			return j == 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p == 1
		}
		res := j >= nums[i] && dfs(i-1, j-nums[i]) || dfs(i-1, j)
		if res {
			*p = 1 // 记忆化
		} else {
			*p = 0
		}
		return res
	}
	return dfs(n-1, s/2)
}

func canPartition2(nums []int) bool {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 != 0 {
		return false
	}

	s /= 2
	n := len(nums)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, s+1)
	}
	f[0][0] = true
	for i, x := range nums {
		for j := 0; j <= s; j++ {
			f[i+1][j] = j >= x && f[i][j-x] || f[i][j]
		}
	}
	return f[n][s]
}

func canPartition3(nums []int) bool {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 != 0 {
		return false
	}

	s /= 2
	f := make([]bool, s+1)
	f[0] = true
	s2 := 0
	for _, x := range nums {
		s2 = min(s2+x, s)
		for j := s2; j >= x; j-- {
			f[j] = f[j] || f[j-x]
		}
		if f[s] {
			return true
		}
	}
	return false
}
