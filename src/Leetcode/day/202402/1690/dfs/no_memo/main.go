package main

func stoneGameVII(stones []int) int {
	n := len(stones)
	s := make([]int, n+1) // 前缀和
	for i, x := range stones {
		s[i+1] = s[i] + x
	}

	// dfs(i,j) 表示剩余石子从 stones[i] 到 stones[j]，先手得分减去后手得分的最大值
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i == j { // 递归边界
			return 0
		}
		// 移除左边或者右边的石子
		return max(s[j+1]-s[i+1]-dfs(i+1, j), s[j]-s[i]-dfs(i, j-1))
	}
	return dfs(0, n-1)
}

func main() {

}
