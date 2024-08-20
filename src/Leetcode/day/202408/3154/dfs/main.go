package dfs

/*
	给你有一个 非负 整数 k 。有一个无限长度的台阶，最低 一层编号为 0 。

	Alice 有一个整数 jump ，一开始值为 0 。Alice 从台阶 1 开始，可以使用 任意 次操作，目标是到达第 k 级台阶。假设 Alice 位于台阶 i ，一次 操作 中，Alice 可以：
	· 向下走一级到 i - 1 ，但该操作 不能 连续使用，如果在台阶第 0 级也不能使用。
	· 向上走到台阶 i + 2^jump 处，然后 jump 变为 jump + 1 。
	请你返回 Alice 到达台阶 k 处的总方案数。

	注意，Alice 可能到达台阶 k 处后，通过一些操作重新回到台阶 k 处，这视为不同的方案。
*/

/*
	需要在递归过程中跟踪以下信息：
	· i：当前位于第 i 个台阶。
	· j：已经使用了 j 次操作二。
	· preDown：一个布尔值，表示上一次操作是否使用了操作一。
	定义状态为 dfs(i,j,preDown)，表示在使用了 j 次操作二，且上一次操作使用/未使用操作一的情况下，从 i 跳到 k 的方案数。

	状态转移方程：
	枚举当前使用哪个操作：
	1. 使用操作二：接下来要解决的问题是 dfs(i+2^j,j+1,false)，将其方案数加到返回值中。
	2. 使用操作一（前提是 preDown=false 且 i>0）：接下来要解决的问题是 dfs(i−1,j,true)，将其方案数加到返回值中。
	3. 此外，如果 i=k（到达终点），则找到了一个方案，把返回值加一。

	注意：到达第 k 个台阶后，还可以继续操作，重新回到第 k 个台阶。所以 i=k 并不是递归边界。
	递归边界：如果 i>k+1，由于操作一不能连续使用，无法到达 k，返回 0。
	递归入口：dfs(1,0,false)，即答案。一开始在第 1 个台阶，没有使用过操作二，也没有使用过操作一。

*/

func waysToReachStair(k int) int {
	type args struct {
		i, j    int
		preDown bool
	}
	memo := make(map[args]int)
	var dfs func(int, int, bool) int
	dfs = func(i int, j int, preDown bool) int {
		if i > k+1 {
			return 0
		}
		p := args{i, j, preDown}
		if v, ok := memo[p]; ok {
			return v
		}
		res := dfs(i+1<<j, j+1, false) // 操作2
		if !preDown && i > 0 {
			res += dfs(i-1, j, true) // 操作 1
		}
		if i == k {
			res++
		}
		memo[p] = res
		return res
	}
	return dfs(1, 0, false)
}
