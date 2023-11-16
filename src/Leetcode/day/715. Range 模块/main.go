package _15__Range_模块

const N = 1e9

type RangeModule struct {
	root *SegmentNode
}

func Constructor() RangeModule {
	return RangeModule{new(SegmentNode)}
}

func (rm *RangeModule) AddRange(left int, right int) {
	rm.root.update(1, N, left, right-1, true)
}

func (rm *RangeModule) QueryRange(left int, right int) bool {
	return rm.root.query(1, N, left, right-1)
}

func (rm *RangeModule) RemoveRange(left int, right int) {
	rm.root.update(1, N, left, right-1, false)
}

type SegmentNode struct {
	lt, rt    *SegmentNode
	val, lazy bool // 覆盖状态和覆盖操作
}

func (t *SegmentNode) update(l, r, i, j int, v bool) {
	if i <= l && r <= j {
		t.val, t.lazy = v, true
		return
	}
	t.pushDown()
	mid := (l + r) >> 1
	if i <= mid {
		t.lt.update(l, mid, i, j, v)
	}
	if j > mid {
		t.rt.update(mid+1, r, i, j, v)
	}
	t.pushUp()
}

func (t *SegmentNode) query(l, r, i, j int) bool {
	if i <= l && r <= j {
		return t.val
	} else if t.val {
		return t.val
	}
	t.pushDown()
	mid, ans := (l+r)>>1, true
	if i <= mid {
		ans = ans && t.lt.query(l, mid, i, j)
	}
	if j > mid {
		ans = ans && t.rt.query(mid+1, r, i, j)
	}
	return ans
}

func (t *SegmentNode) pushUp() { t.val = t.lt.val && t.rt.val }

func (t *SegmentNode) pushDown() {
	if t.lt == nil {
		t.lt = new(SegmentNode)
	}
	if t.rt == nil {
		t.rt = new(SegmentNode)
	}
	if t.lazy {
		t.lt.val, t.lt.lazy = t.val, true
		t.rt.val, t.rt.lazy = t.val, true
		t.lazy = false
	}
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
