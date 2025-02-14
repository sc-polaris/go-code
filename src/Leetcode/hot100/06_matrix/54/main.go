package main

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func spiralOrder(matrix [][]int) (ans []int) {
	m, n := len(matrix), len(matrix[0])
	k := m * n
	i, j := 0, -1
	for di := 0; len(ans) < k; di = (di + 1) % 4 {
		for range n {
			i += dirs[di][0]
			j += dirs[di][1]
			ans = append(ans, matrix[i][j])
		}
		n, m = m-1, n
	}
	return
}
