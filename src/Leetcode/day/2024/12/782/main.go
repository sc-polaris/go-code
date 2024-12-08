package main

/*
	一个 n x n 的二维网络 board 仅由 0 和 1 组成 。每次移动，你能交换任意两列或是两行的位置。

	返回 将这个矩阵变为  “棋盘”  所需的最小移动次数 。如果不存在可行的变换，输出 -1。

	“棋盘” 是指任意一格的上下左右四个方向的值均与本身不同的矩阵。
*/

/*
	board 要能变成棋盘，有两个必要条件：
	1. 只有两种行：A 行是从 0101⋯ 交换得到的，也就是说，A 行中的 0 的个数和 1 的个数相差至多为 1；B 行是从 1010⋯ 交换得到的，和 A 行完全相反。
	2. 这两种行的个数相差至多为 1。

	不能变成棋盘的情况
	核心思路：统计 A 行和 B 行的个数。统计的过程中，如果发现有一行既不是 A 行也不是 B 行，直接返回 −1。统计结束后，如果发现 A 行和 B 行的个数相差超过 1，也返回 −1。
	可以用哈希表统计。但是，没有这个必要。
	以 board 第一行（或者任意一行）为参照物。
	统计 board 第一行中的 0 和 1 的个数，如果 0 和 1 的个数相差超过 1，说明这一行既不是 A 行也不是 B 行，无论如何交换，绝不可能得到棋盘，返回 −1。
	对于其余行 board[i]，首先比较 board[i][0] 和 board[0][0]：
	· 相同：那么 board[i] 必须和 board[0] 完全相同，即对于任意 j，board[i][j]==board[0][j] 成立。若不满足，返回 −1。
	· 不同：那么 board[i] 必须和 board[0] 完全不同，即对于任意 j，board[i][j]!=board[0][j] 成立。若不满足，返回 −1。
	如果没有返回 −1，那么说明只有 A 行和 B 行。我们也只需统计 board[i][0]（第一列）的 0 和 1 的个数，就知道有多少个 A 行和 B 行。如果 0 和 1 的个数相差超过 1，那么返回 −1。（代码实现时，关于第一列的判断可以放在前面）

	如果上述情况都没有返回 −1，那么 board 一定可以变成棋盘，方法如下。

	最小交换次数
	比如把 s=001110 通过交换元素，变成 t=010101。这其中 s[i]!=t[i] 出现了 4 次。我们需要让 s 中的 0 在偶数下标上，1 在奇数下标上。那么把奇数上的 0 和偶数上的 1 交换，
	就可以满足要求，所以只需要 4/2=2 次交换。一般地，有如下定理。
*/

func movesToChessboard(board [][]int) int {
	n := len(board)
	firstRow := board[0]
	firstCol := make([]int, n)
	var rowCnt, colCnt [2]int
	for i, row := range board {
		rowCnt[firstRow[i]]++ // 统计 0 和 1 的个数
		firstCol[i] = row[0]
		colCnt[firstCol[i]]++
	}

	// 第一行，0 和 1 的个数之差不能超过 1
	// 第一列，0 和 1 的个数之差不能超过 1
	if abs(rowCnt[0]-rowCnt[1]) > 1 || abs(colCnt[0]-colCnt[1]) > 1 {
		return -1
	}

	// 每一行和第一行比较，要么完全相同，要么完全不同
	for _, row := range board {
		same := row[0] == firstRow[0]
		for i, x := range row {
			if (x == firstRow[i]) != same {
				return -1
			}
		}
	}

	return minSwap(firstRow, rowCnt) + minSwap(firstCol, colCnt)
}

// 计算最小交换次数
func minSwap(arr []int, cnt [2]int) int {
	x0 := 0 // 如果 n 是偶数，x0 是 0
	if cnt[1] > cnt[0] {
		x0 = 1
	}
	diff := 0
	for i, x := range arr {
		diff += i%2 ^ x ^ x0
	}
	n := len(arr)
	if n%2 > 0 {
		return diff / 2
	}
	return min(diff, n-diff) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
