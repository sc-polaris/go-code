package main

/*
	在二维平面上的 x 轴上，放置着一些方块。

	给你一个二维整数数组 positions ，其中 positions[i] = [lefti, sideLengthi] 表示：第 i 个方块边长为 sideLengthi ，其左侧边与 x 轴上坐标点 lefti 对齐。

	每个方块都从一个比目前所有的落地方块更高的高度掉落而下。方块沿 y 轴负方向下落，直到着陆到 另一个正方形的顶边 或者是 x 轴上 。一个方块仅仅是擦过另一个方块的左侧
	边或右侧边不算着陆。一旦着陆，它就会固定在原地，无法移动。

	在每个方块掉落后，你必须记录目前所有已经落稳的 方块堆叠的最高高度 。

	返回一个整数数组 ans ，其中 ans[i] 表示在第 i 块方块掉落后堆叠的最高高度。
*/

// fallingSquares1 暴力
func fallingSquares1(positions [][]int) []int {
	n := len(positions)
	h := make([]int, n)
	for i, p := range positions {
		l1, r1 := p[0], p[0]+p[1]-1
		h[i] = p[1]
		for j, q := range positions[:i] {
			l2, r2 := q[0], q[0]+q[1]-1
			if r1 >= l2 && r2 >= l1 {
				h[i] = max(h[i], h[j]+p[1])
			}
		}
	}
	for i := 1; i < n; i++ {
		h[i] = max(h[i], h[i-1])
	}
	return h
}

/*
	方法二：线段树
	根据题目描述，我们需要维护一个区间集合，支持区间的修改和查询操作。这种情况下，我们可以使用线段树来解决。

	线段树将整个区间分割为多个不连续的子区间，子区间的数量不超过 log(width)，其中 width 是区间的长度。更新某个元素的值，
	只需要更新 log(width) 个区间，并且这些区间都包含在一个包含该元素的大区间内。区间修改时，需要使用懒标记保证效率。
	· 线段树的每个节点代表一个区间；
	· 线段树具有唯一的根节点，代表的区间是整个统计范围，如 [1,n]；
	· 线段树的每个叶子节点代表一个长度为 1 的元区间 [x,x]；
	· 对于每个内部节点 [l,r]，它的左儿子是 [l,mid]，右儿子是 [mid+1,r], 其中 mid=(l+r)/2；

	对于本题，线段树节点维护的信息有：
	1. 区间中方块的最大高度 v
	2. 懒标记 add
	另外，由于数轴范围很大，达到 10^8，因此我们采用动态开点。
*/

type node struct {
	left      *node
	right     *node
	l, mid, r int
	v, add    int
}

func newNode(l, r int) *node {
	return &node{l: l, r: r, mid: (l + r) >> 1}
}

type segmentTree struct {
	root *node
}

func NewSegmentTree() *segmentTree {
	return &segmentTree{root: newNode(1, 1e9)}
}

func (t *segmentTree) modify(l, r, v int, n *node) {
	if l > r {
		return
	}
	if n.l >= l && n.r <= r {
		n.v = v
		n.add = v
		return
	}
	t.pushDown(n)
	if l <= n.mid {
		t.modify(l, r, v, n.left)
	}
	if r > n.mid {
		t.modify(l, r, v, n.right)
	}
	t.pushUp(n)
}

func (t *segmentTree) query(l, r int, n *node) int {
	if l > r {
		return 0
	}
	if n.l >= l && n.r <= r {
		return n.v
	}
	t.pushDown(n)
	v := 0
	if l <= n.mid {
		v = max(v, t.query(l, r, n.left))
	}
	if r > n.mid {
		v = max(v, t.query(l, r, n.right))
	}
	return v
}

func (t *segmentTree) pushUp(n *node) {
	n.v = max(n.left.v, n.right.v)
}

func (t *segmentTree) pushDown(n *node) {
	if n.left == nil {
		n.left = newNode(n.l, n.mid)
	}
	if n.right == nil {
		n.right = newNode(n.mid+1, n.r)
	}
	if n.add != 0 {
		n.left.add = n.add
		n.right.add = n.add
		n.left.v = n.add
		n.right.v = n.add
		n.add = 0
	}
}

func fallingSquares(positions [][]int) []int {
	ans := make([]int, len(positions))
	t := NewSegmentTree()
	mx := 0
	for i, p := range positions {
		l, w, r := p[0], p[1], p[0]+p[1]-1
		h := t.query(l, r, t.root) + w
		mx = max(mx, h)
		ans[i] = mx
		t.modify(l, r, h, t.root)
	}
	return ans
}
