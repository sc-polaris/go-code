package main

/*

	状态定义：定义 f[i][j] 表示切割一块高 i 宽 j 的木块，能得到的最多钱数。
	分类讨论：
		1. 如果直接售卖，则收益对应的 price（如果存在的化）
		2. 如果竖着切开，枚举切割位置（宽度）k，得到两个高为 i，宽分别为 k 和 j-k 的木块，最大收益为：
			max f[i][k] + f[i][j-k]		k=1 j-1
		3. 如果横着切开，枚举切割位置（高度）k，得到两个宽为 j，高分别为 k 和 i-k 的木块，最大收益为：
			max f[k][j] + f[i-k][j]		k=1 i-1
		取上述三种情况的最大值，即为 f[i][j]

*/

// sellingWood1 未优化
func sellingWood1(m int, n int, prices [][]int) int64 {
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for _, price := range prices {
		pr[price[0]][price[1]] = price[2]
	}

	f := make([][]int64, m+1)
	for i := 1; i <= m; i++ {
		f[i] = make([]int64, n+1)
		for j := 1; j <= n; j++ {
			f[i][j] = int64(pr[i][j])
			for k := 1; k < j; k++ { // 垂直切割，枚举宽度 k
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k < i; k++ { // 水平切割，枚举高度 k
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return f[m][n]
}

// sellingWood 优化数组 pr 没必要存在
func sellingWood(m int, n int, prices [][]int) int64 {
	f := make([][]int64, m+1)
	for i := range f {
		f[i] = make([]int64, n+1)
	}
	for _, price := range prices {
		f[price[0]][price[1]] = int64(price[2])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= j/2; k++ { // 垂直切割，枚举宽度 k
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k <= i/2; k++ { // 水平切割，枚举高度 k
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return f[m][n]
}

func main() {

}
