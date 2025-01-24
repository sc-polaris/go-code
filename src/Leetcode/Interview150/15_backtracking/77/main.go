package main

/*
	方法一：枚举下一个数选哪个
*/

func combine(n int, k int) (ans [][]int) {
	var path []int
	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path) // 还要选 d 个数
		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j - 1)
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(n)
	return
}

/*
	方法二：选或不选
*/

func combine2(n, k int) (ans [][]int) {
	var path []int
	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path) // 还要选 d 个数
		if d == 0 {        // 选好了
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 如果 i > d，可以不选 i
		if i > d {
			dfs(i - 1)
		}

		path = append(path, i)
		dfs(i - 1)
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(n)
	return
}
