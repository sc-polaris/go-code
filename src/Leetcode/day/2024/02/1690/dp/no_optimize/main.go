package main

func stoneGameVII(stones []int) int {
	n := len(stones)
	s := make([]int, n+1)
	for i, x := range stones {
		s[i+1] = s[i] + x
	}
	f := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			// 移除左边或者右边的石子
			f[i][j] = max(s[j+1]-s[i+1]-f[i+1][j], s[j]-s[i]-f[i][j-1])
		}
	}
	return f[0][n-1]
}

func main() {

}
