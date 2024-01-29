package main

func findRotateSteps(s string, t string) int {
	n, m := len(s), len(t)
	st := make([][]bool, n)
	for i := range st {
		st[i] = make([]bool, m+1)
	}
	st[0][0] = true
	type pair struct{ i, j int }
	q := []pair{{0, 0}}
	for step := 0; ; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			i, j := p.i, p.j
			if j == m {
				return step
			}
			// 移动到 (i,j+1)
			if s[i] == t[j] {
				if st[i][j+1] {
					continue
				}
				st[i][j+1] = true
				q = append(q, pair{i, j + 1})
			}
			// 否则向左移动 或者向右一动
			for _, i2 := range []int{(i - 1 + n) % n, (i + 1) % n} {
				if st[i2][j] {
					continue
				}
				st[i2][j] = true
				q = append(q, pair{i2, j})
			}
		}
	}
}

func main() {

}
