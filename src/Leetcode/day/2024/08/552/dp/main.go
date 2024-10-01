package dp

// f[i][j][k]「在之前填过 j 个 A，且右边相邻位置有 k 个连续 L 的情况下，继续填字母，能构造多少个长为 i 的字符串」。

func checkRecord(n int) int {
	const mod = 1_000_000_007
	const mx = 100_001
	var f [mx][2][3]int
	f[0] = [2][3]int{
		{1, 1, 1},
		{1, 1, 1},
	}
	for i := 1; i < mx; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 3; k++ {
				res := f[i-1][j][0]
				if j == 0 {
					res += f[i-1][1][0]
				}
				if k < 2 {
					res += f[i-1][j][k+1]
				}
				f[i][j][k] = res % mod
			}
		}
	}
	return f[n][0][0]
}
