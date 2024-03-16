package main

/*
	一开始把所以 (i,0) 都加入一个列表，遍历列表中的坐标。
	只有之前没入队的格子才能入队，用数组标记
*/

func maxMoves1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	st := make([]int, m)
	q := make([]int, m)
	for i := range q {
		q[i] = i
	}
	for j := 0; j < n-1; j++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if st[k] != j+1 && grid[k][j+1] > grid[i][j] {
					st[k] = j + 1 // 第 k 行目前最右访问到了 j
					q = append(q, k)
				}
			}
		}
		if q == nil { // 无法再往右边走了
			return j
		}
	}
	return n - 1
}

// maxMoves 优化
func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for _, row := range grid {
		row[0] *= -1 // 入队标记
	}
	for j := 0; j < n-1; j++ {
		ok := false
		for i := 0; i < m; i++ {
			if grid[i][j] > 0 { // 不在队列中
				continue
			}
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if grid[k][j+1] > -grid[i][j] {
					grid[k][j+1] *= -1 // 入队标记
					ok = true
				}
			}
		}
		if !ok { // 无法再往右走了
			return j
		}
	}
	return n - 1
}

func main() {

}
