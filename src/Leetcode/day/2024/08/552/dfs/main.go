package dfs

/*
	可以用字符串表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
	· 'A'：Absent，缺勤
	· 'L'：Late，迟到
	· 'P'：Present，到场

	如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
	· 按 总出勤 计，学生缺勤（'A'）严格 少于两天。
	· 学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。

	给你一个整数 n ，表示出勤记录的长度（次数）。请你返回记录长度为 n 时，可能获得出勤奖励的记录情况 数量 。答案可能很大，所以返回对 109 + 7 取余 的结果。

	dfs(i,j,k)「在之前填过 j 个 A，且右边相邻位置有 k 个连续 L 的情况下，继续填字母，能构造多少个长为 i 的字符串」。
	考虑长为 i 的字符串的最后一个位置填什么字母：
	· 填 P：问题变成，在之前填过 j 个 A，且右边相邻位置有 0 个连续 L 的情况下，继续填字母，能构造多少个长为 i−1 的字符串，即 dfs(i−1,j,0)。
	· 如果 j=0，那么可以填 A：问题变成，在之前填过 1 个 A，且右边相邻位置有 0 个连续 L 的情况下，继续填字母，能构造多少个长为 i−1 的字符串，即 dfs(i−1,1,0)。
	· 如果 k<2，那么可以填 L：问题变成，在之前填过 j 个 A，且右边相邻位置有 k+1 个连续 L 的情况下，继续填字母，能构造多少个长为 i−1 的字符串，即 dfs(i−1,j,k+1)。
*/

func checkRecord(n int) int {
	const mod = 1_000_000_007
	const mx = 100_001
	var memo [mx][2][3]int
	var dfs func(int, int, int) int
	dfs = func(i int, j int, k int) int {
		if i == 0 {
			return 1
		}
		p := &memo[i][j][k]
		if *p > 0 {
			return *p
		}
		res := dfs(i-1, j, 0)
		if j == 0 {
			res += dfs(i-1, 1, 0)
		}
		if k < 2 {
			res += dfs(i-1, j, k+1)
		}
		*p = res % mod
		return *p
	}
	return dfs(n, 0, 0)
}
