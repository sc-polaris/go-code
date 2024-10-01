package main

/*
去除递归中的「递」，只保留「归」的部分，即自底向上计算
*/

func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	a := make([]int, m) // a[i] 是第 i 排可用椅子的下标集合
	for i, s := range seats {
		for j, c := range s {
			if c == '.' {
				a[i] |= 1 << j
			}
		}
	}

	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, 1<<n)
	}
	for j := 1; j < 1<<n; j++ {
		lb := j & -j
		f[0][j] = f[0][j&^(lb*3)] + 1
	}
	for i := 1; i < m; i++ {
		for j := a[i]; j > 0; j = (j - 1) & a[i] { // 枚举 a[i] 的子集 j
			f[i][j] = f[i-1][a[i-1]]             // 第 i 排空这
			for s := j; s > 0; s = (s - 1) & j { // 枚举 j 的子集 s
				if s&(s>>1) == 0 { // s 没有连续的 1
					t := a[i-1] &^ (s<<1 | s>>1) // 去掉不能坐人的位置
					f[i][j] = max(f[i][j], f[i-1][t]+f[0][s])
				}
			}
		}
		f[i][0] = f[i-1][a[i-1]]
	}
	return f[m-1][a[m-1]]
}

func main() {

}
