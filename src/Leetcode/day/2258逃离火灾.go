package day

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 1. 二分
//func maximumMinutes(grid [][]int) int {
//	m, n := len(grid), len(grid[0])
//	ans := sort.Search(m*n+1, func(t int) bool {
//		onFire := make([][]bool, m)
//		for i := range onFire {
//			onFire[i] = make([]bool, n)
//		}
//		var f []pair
//		for i, row := range grid {
//			for j, x := range row {
//				if x == 1 {
//					onFire[i][j] = true // 标记着火的位置
//					f = append(f, pair{i, j})
//				}
//			}
//		}
//		// 火的bfs
//		spreadFire := func() {
//			tmp := f
//			f = nil
//			for _, p := range tmp {
//				for _, d := range dirs { // 枚举上下左右四个方向
//					x, y := p.x+d.x, p.y+d.y
//					if 0 <= x && x < m && 0 <= y && y < n && !onFire[x][y] && grid[x][y] == 0 {
//						onFire[x][y] = true // 标记着火的位置
//						f = append(f, pair{x, y})
//					}
//				}
//			}
//		}
//		for ; t > 0 && len(f) > 0; t-- { // 如果火无法扩散就提前推出
//			spreadFire() // 火扩散
//		}
//		if onFire[0][0] {
//			return true // 起点着火 g
//		}
//
//		// 人的bfs
//		vis := make([][]bool, m)
//		for i := range vis {
//			vis[i] = make([]bool, n)
//		}
//		vis[0][0] = true
//		q := []pair{{0, 0}}
//		for len(q) > 0 {
//			tmp := q
//			q = nil
//			for _, p := range tmp {
//				if onFire[p.x][p.y] { // 人走到这个位置，火也扩散到了这个位置
//					continue
//				}
//				for _, d := range dirs { // 枚举上下左右四个方向
//					x, y := p.x+d.x, p.y+d.y
//					if 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && !onFire[x][y] && grid[x][y] == 0 {
//						if x == m-1 && y == n-1 {
//							return false // 安全了
//						}
//						vis[x][y] = true // 避免重复访问同一个位置
//						q = append(q, pair{x, y})
//					}
//				}
//			}
//			spreadFire() // 火蔓延
//		}
//
//		return true // 人被烧到，或者没有可以到达安全屋的路
//	}) - 1
//	if ans < m*n {
//		return ans
//	}
//	return 1_000_000_000
//}

// 2. 直接算
func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 返回三个数，分别表示到达安全屋/安全屋左边/安全屋上边的最短时间
	bfs := func(q []pair) (int, int, int) {
		time := make([][]int, m)
		for i := range time {
			time[i] = make([]int, n)
			for j := range time[i] {
				time[i][j] = -1 // -1 表示未访问
			}
		}
		for _, p := range q {
			time[p.x][p.y] = 0
		}
		for t := 1; len(q) > 0; t++ { // 每次循环向外扩展一圈
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dirs {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 0 && time[x][y] < 0 {
						time[x][y] = t
						q = append(q, pair{x, y})
					}
				}
			}
		}
		return time[m-1][n-1], time[m-1][n-2], time[m-2][n-1]
	}

	manToHouseTime, m1, m2 := bfs([]pair{{0, 0}})
	if manToHouseTime < 0 { // 人无法到达安全屋
		return -1
	}

	firePos := []pair{}
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				firePos = append(firePos, pair{i, j})
			}
		}
	}
	fireToHouseTime, f1, f2 := bfs(firePos) // 火哥着火点同时跑bfs
	if fireToHouseTime < 0 {                // 火无法到达安全屋
		return 1_000_000_000
	}

	d := fireToHouseTime - manToHouseTime
	if d < 0 { // 火比人先到安全屋
		return -1
	}
	if m1 != -1 && m1+d < f1 || // 安全屋左边相邻格子，人比火先到
		m2 != -1 && m2+d < f2 { // 安全屋上边相邻格子，人比火先到
		return d
	}

	return d - 1
}
