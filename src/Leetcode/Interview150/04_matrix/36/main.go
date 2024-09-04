package main

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]bool
	var box [3][3][9]bool
	for i, row := range board {
		for j, col := range row {
			if col == '.' {
				continue
			}
			index := int(col - '1')
			if rows[i][index] || columns[j][index] || box[i/3][j/3][index] {
				return false
			}
			rows[i][index] = true
			columns[j][index] = true
			box[i/3][j/3][index] = true
		}
	}
	return true
}
