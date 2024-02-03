package main

func stoneGameVII(stones []int) int {
	n := len(stones)
	s := make([]int, n+1)
	for i, x := range stones {
		s[i+1] = s[i] + x
	}
	f := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 移除左边或者右边的石子
			f[j] = max(s[j+1]-s[i+1]-f[j], s[j]-s[i]-f[j-1])
		}
	}
	return f[n-1]
}

func main() {

}
