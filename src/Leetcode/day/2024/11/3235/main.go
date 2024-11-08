package main

/*
	给你两个正整数 xCorner 和 yCorner 和一个二维整数数组 circles ，其中 circles[i] = [xi, yi, ri] 表示一个圆心在 (xi, yi) 半径为 ri 的圆。

	坐标平面内有一个左下角在原点，右上角在 (xCorner, yCorner) 的矩形。你需要判断是否存在一条从左下角到右上角的路径满足：路径 完全 在矩形内部，不会
	触碰或者经过 任何 圆的内部和边界，同时 只 在起点和终点接触到矩形。

	如果存在这样的路径，请你返回 true ，否则返回 false 。
*/

/*
	如果从矩形【上边界/左边界】到矩形【右边界/下边界】的路被圆堵死，则无法从矩形左下角移动到矩形右上角。

	怎么判断呢？

	情况一：两圆交集在矩阵外
	只考虑圆和矩形相交相切的情况
	这两个圆都是 DFS 的起点/终点。
	看上去连边不影响答案
	但考虑存在下图的情况
	连边会算错，所以应该不连边。

	情况二：两圆交集在矩阵内
	说明圆和矩形相交
	这两个圆都是 DFS 的起点/终点。
	如果出现上图的情况
	圆必定包含矩形右上角，
	两个节点是否连边不影响答案。

	情况三：两圆交集完全在矩形内。一定要在两个节点之间连边

	结论：两圆交集中可以任选一个点，如果其严格在矩形内，则连边，否则不连

	具体做法
	从与矩形【上边界/左边界】相交/相切的圆开始 DFS。

	如果当前 DFS 到了圆 i：
	· 先判断其是否与矩形【右边界/下边界】相交或相切，如果是，则 DFS 返回 true。
	· 否则，判断其是否与其他圆 j 相交或相切，如果是，则判断点 A 是否严格在矩形内，如果在，则递归 j，如果收到了 true，则 DFS 返回 true。
	最后，如果最外层调用 DFS 的地方收到了 true，则表示无法从矩形左下角移动到矩形右上角，返回 false。

	代码实现时，可以在递归之前，特判圆包含矩形左下角或者矩形右上角的情况，此时可以直接返回 false。

*/

// 判断点 (x,y) 是否在圆 (ox,oy,r) 内
func inCircle(ox, oy, r, x, y int) bool {
	return (ox-x)*(ox-x)+(oy-y)*(oy-y) <= r*r
}

func canReachCorner(X int, Y int, circles [][]int) bool {
	vis := make([]bool, len(circles))
	var dfs func(int) bool
	dfs = func(i int) bool {
		x1, y1, r1 := circles[i][0], circles[i][1], circles[i][2]
		// 圆 i 是否与矩形右边界/下边界相交相切
		if y1 <= Y && abs(x1-X) <= r1 || x1 <= X && y1 <= r1 || x1 > X && inCircle(x1, y1, r1, X, 0) {
			return true
		}
		vis[i] = true
		for j, c := range circles {
			x2, y2, r2 := c[0], c[1], c[2]
			// 在两圆相交相切的前提下，点 A 是否严格在矩形内
			if !vis[j] && (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) <= (r1+r2)*(r1+r2) &&
				x1*r2+x2*r1 < (r1+r2)*X &&
				y1*r2+y2*r1 < (r1+r2)*Y &&
				dfs(j) {
				return true
			}
		}
		return false
	}
	for i, c := range circles {
		x, y, r := c[0], c[1], c[2]
		if inCircle(x, y, r, 0, 0) || // 圆 i 包含矩形左下角
			inCircle(x, y, r, X, Y) || // // 圆 i 包含矩形右上角
			// 圆 i 是否与矩形上边界/左边界相交相切
			!vis[i] && (x <= X && abs(y-Y) <= r || y <= Y && x <= r || y > Y && inCircle(x, y, r, 0, Y)) && dfs(i) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
