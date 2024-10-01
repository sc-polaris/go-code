package main

/*
	树的直径
	以示例 1 为例子，我们可以把问题拆分成两个问题。

	第一个问题：以 3 为根节点的子树的 二叉树的最大深度，其中节点 3 的深度为 0.
	这棵子树的最大深度是 111，也就是 333 往下走的最远距离

	第二个问题：去掉 3 的所有子孙节点，此时 3 变成了一个叶子节点。我们相当于计算某一个端点固定为 3 的 二叉树的直径。
	若固定直径的一个端点为 3，则这条直径可以是 3-1-5-4-9，长度（边的个数）为 4。相当于从 3 出发往上走，在某个点「拐弯」后再往下走

	这两个问题取最大值 max(1,4) = 4，即为原问题的答案。

	本问题的算法：
	1. 递归时，除了返回当前子树的最大链长加一，还需要返回一个布尔值，表示当前子树是否包含 start。
	2. 如果当前节点是空节点，返回 0 和 false。
	3. 设左子树的返回的链长为 lLen，右子树返回的链长为 rLen。
	4. 如果当前节点值等于 start，初始化答案为 max(lLen,rLen)，即 start 的最大深度。然后返回 1 和 true。
	5. 如果左右子树都不包含 start，返回 max(lLen,rLen) + 1。
	6. 如果左子树或右子树包含 start，像计算直径那样，用 lLen + rLen 更新答案的最大值。如果左子树包含 start，则返回 lLen 和 true，
	   否则返回 rLen 和 true。这种返回方式可以保证 lLen + rLen 一定是端点为 start 的直径长度。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) (ans int) {
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, false
		}
		lLen, lFound := dfs(node.Left)
		rLen, rFound := dfs(node.Right)
		if node.Val == start {
			// 计算子树 start 的最大深度
			// 注意这里和方法一的区别，max 后面没有 +1，所以算出的也是最大深度。
			ans = max(lLen, rLen)
			return 1, true // 找到了 start
		}
		if lFound || rFound {
			// 只有在左子树或右子树包含 start 时，才能更新答案
			ans = max(ans, lLen+rLen) // 两条链拼成直径
			// 保证 start 是直径端点
			if lFound {
				return lLen + 1, true
			}
			return rLen + 1, true
		}
		return max(lLen, rLen) + 1, false
	}
	dfs(root)
	return
}

/*
	上面的写法 DFS 会返回两个值，能否只返回一个值呢？
	去掉布尔值，把负数利用起来：
	· 如果子树不包含 start，将原来的返回值改成相反数（负数）。
	· 如果子树包含 start，返回值不变。
*/

func amountOfTime2(root *TreeNode, start int) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lLen := dfs(node.Left)
		rLen := dfs(node.Right)
		if node.Val == start {
			ans = -min(lLen, rLen) // 负负得正
			return 1               // 用正数表示找到了 start
		}
		if lLen > 0 || rLen > 0 {
			ans = max(ans, abs(lLen)+abs(rLen))
			return max(lLen, rLen) + 1 // max 会自动取到正数
		}
		return min(lLen, rLen) - 1 // 用负数表示没有找到 start
	}
	dfs(root)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
