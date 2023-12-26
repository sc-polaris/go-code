package main

func maxStudents(seats [][]byte) int {
	m := len(seats)
	a := make([]int, m)
	for i, s := range seats {
		for j, c := range s {
			if c == '.' {
				a[i] |= 1 << j
			}
		}
	}

	memo := map[[3]int]int{}
	var dfs func(int, int, int) int
	dfs = func(i int, j int, k int) (res int) {
		t := [3]int{i, j, k}
		if v, ok := memo[t]; ok {
			return v
		}
		defer func() { memo[t] = res }()
		if j == 0 {
			if i == 0 {
				return 0
			}
			return dfs(i-1, a[i-1]&^(k<<1|k>>1), 0)
		}
		lb := j & -j
		return max(dfs(i, j^lb, k), dfs(i, j&^(lb*3), k|lb)+1)
	}
	return dfs(m-1, a[m-1], 0)
}

func main() {

}
