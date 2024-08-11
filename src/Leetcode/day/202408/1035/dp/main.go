package dp

/*
	定义 dfs(i,j) 表示 s[0] 到 s[i] 与 t[0] 到 t[j] 之间的最大连线数。
	状态转移方程
						dfs(i-1,j-1) + 1			s[i] == t[j]
			dfs(i,j) =
						max(dfs(i-1,j),dfs(i,j-1))	s[i] != t[j]

	定义 f[i+1][j+1] 表示 s[0] 到 s[i] 与 t[0] 到 t[j] 之间的最大连线数。
	状态转移方程
								f[i][j] + 1					s[i] == t[j]
				f[i+1][j+1] =
								max(f[i][j+1],f[i+1][j])	s[i] != t[j]
*/

func maxUncrossedLines(s []int, t []int) int {
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i, x := range s {
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}

/*
	观察上面的状态转移方程，在计算 f[i+1] 时，只会用到 f[i]，不会用到比 i 更早的状态。
	空间优化 去掉第一个维度，把 f[i+1] 和 f[i] 保存到同一个数组中。
	状态转移方程改为

					f[j] + 1		s[i] == t[j]
		f[j+1] =
					max(f[j+1],f[j])	s[i] != t[j]

	当 s[i]=t[j] 时，计算 f[j+1] 会用到 f[j]（这里相当于空间优化前的 f[i][j]），但 f[j] 已经被覆盖成 f[i+1][j] 了，怎么办？
	用一个变量 pre 记录被覆盖前的 f[j] 即可。

	注意 s[i]!=t[j] 时，转移来源 f[j] 是空间优化前的 f[i+1][j]，我们需要的正是被覆盖的 f[j]，所以这里是不需要 pre 的。
*/

func maxUncrossedLines2(s []int, t []int) int {
	m := len(t)
	f := make([]int, m+1)
	for _, x := range s {
		pre := 0 // f[0]
		for j, y := range t {
			if x == y {
				f[j+1], pre = pre+1, f[j+1]
			} else {
				pre = f[j+1]
				f[j+1] = max(f[j+1], f[j])
			}
		}
	}
	return f[m]
}
