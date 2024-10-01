package main

/*
	dfs(i,j) 表示当前 jug1 中有 i 升水，jug2 中有 j 升水，是否可以得到 z 升水、
	执行过程：
	1. 如果 (i,j) 已经被访问过，返回 false
	2. 如果 i=z 或者 j=z 或者 i+j=z，返回 true
	3. 如果我们给 jug1 倒满水，或者给 jug2 倒满水，或者将 jdg1 清空，或者将 jdg2 清空，可以得到 z 升水，返回 true
	4. 如果我们将 jug1 中的水导入 jug2，或者将 jdg2 中的水倒入 jdg1，可以得到 z 升水，返回 true
*/

func canMeasureWater(x int, y int, z int) bool {
	type pair struct{ x, y int }
	st := make(map[pair]bool)
	var dfs func(int, int) bool
	dfs = func(i int, j int) bool {
		p := pair{i, j}
		if st[p] {
			return false
		}
		st[p] = true
		if i == z || j == z || i+j == z {
			return true
		}
		if dfs(x, j) || dfs(i, y) || dfs(0, j) || dfs(i, 0) {
			return true
		}
		a := min(i, y-j)
		b := min(j, x-i)
		return dfs(i-a, j+a) || dfs(i+b, j-b)
	}
	return dfs(0, 0)
}

func main() {

}
