package main

func rob(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		*p = max(dfs(i-1), dfs(i-2)+nums[i])
		return *p
	}
	return dfs(n - 1)
}

func rob2(nums []int) int {
	n := len(nums)
	f := make([]int, n+2)
	for i, x := range nums {
		f[i+2] = max(f[i+1], f[i]+x)
	}
	return f[n+1]
}

func rob3(nums []int) int {
	f0, f1 := 0, 0
	for _, x := range nums {
		f0, f1 = f1, max(f1, f0+x)
	}
	return f1
}
