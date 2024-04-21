package main

/*
	选 or 不选
	剪枝：
	最大 d 个数的和
	i + i-1 + i-2 + ... + i-d+1 = (i+i-d+1)*d/2 = (i*2-d+1)*d/2
*/

func combinationSum3(k int, n int) (ans [][]int) {
	var path []int
	var dfs func(int, int)
	dfs = func(i, t int) {
		d := k - len(path) // 还要选 d 个数
		if t < 0 || t > (i*2-d+1)*d/2 {
			return
		}
		if d == 0 { // 找到一个合法组合
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 不选 i
		if i > d {
			dfs(i-1, t)
		}

		// 选 i
		path = append(path, i)
		dfs(i-1, t-i)
		path = path[:len(path)-1] // 回溯
	}
	dfs(9, n)
	return
}
