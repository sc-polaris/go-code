package main

import "math/bits"

// 优化的dfs
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

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, 1<<n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没计算过
		}
	}

	var dfs func(int, int) int
	dfs = func(i int, j int) (res int) {
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }()
		if i == 0 {
			if j == 0 {
				return 0
			}
			lb := j & -j
			return dfs(i, j&^(lb*3)) + 1
		}
		res = dfs(i-1, a[i-1])               // 第 i 排空着
		for s := j; s > 0; s = (s - 1) & j { // 枚举 j 的子集 s
			if s&(s>>1) == 0 { // s 没有连续的1
				t := a[i-1] &^ (s<<1 | s>>1) // 去掉不能坐人的位置
				res = max(res, dfs(i-1, t)+bits.OnesCount(uint(s)))
			}
		}
		return res
	}

	return dfs(m-1, a[m-1])
}

func main() {

}
