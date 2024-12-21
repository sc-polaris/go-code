package main

import "slices"

/*
	班里有 m 位学生，共计划组织 n 场考试。给你一个下标从 0 开始、大小为 m x n 的整数矩阵 score ，其中每一行对应一位学生，而 score[i][j] 表示第 i 位学生
	在第 j 场考试取得的分数。矩阵 score 包含的整数 互不相同 。

	另给你一个整数 k 。请你按第 k 场考试分数从高到低完成对这些学生（矩阵中的行）的排序。

	返回排序后的矩阵。
*/

func sortTheStudents(score [][]int, k int) [][]int {
	slices.SortFunc(score, func(a, b []int) int { return b[k] - a[k] })
	return score
}
