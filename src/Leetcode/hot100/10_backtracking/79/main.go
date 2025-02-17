package main

import "slices"

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	cnt := make(map[byte]int)
	for _, row := range board {
		for _, c := range row {
			cnt[c]++
		}
	}
	w := []byte(word)
	wordCnt := make(map[byte]int)
	for _, c := range w {
		wordCnt[c]++
		if wordCnt[c] > cnt[c] {
			return false
		}
	}

	if cnt[w[len(w)-1]] < cnt[w[0]] {
		slices.Reverse(w)
	}

	m, n := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		board[i][j] = '0' // 标记访问过
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n && dfs(x, y, k+1) {
				return true
			}
		}
		board[i][j] = word[k] // 恢复
		return false
	}
	for i := range m {
		for j := range n {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
