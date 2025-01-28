package main

/*
	给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
	在「杨辉三角」中，每个数是它左上方和右上方的数的和。
*/

// 组合数学公式
func getRow2(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}

// 预处理
var c [34][]int

func init() {
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0], c[i][i] = 1, 1
		for j := i; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

func getRow(rowIndex int) []int {
	return c[rowIndex]
}
