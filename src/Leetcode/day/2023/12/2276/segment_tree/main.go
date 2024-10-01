package main

// 动态开点线段树

type CountIntervals struct {
	left, right *CountIntervals
	l, r, cnt   int // 线段树的每个节点可以保存对应范围的左右端点 l 和 r，以及范围内 add 过的整数个数 cnt。
}

func Constructor() CountIntervals { return CountIntervals{l: 1, r: 1e9} }

func (o *CountIntervals) Add(l int, r int) {
	if o.cnt == o.r-o.l+1 { // o 已被完整覆盖，无需执行任何操作
		return
	}
	if l <= o.l && o.r <= r { // 当前节点已被区间 [l,r] 完整覆盖，不再继续递归
		o.cnt = o.r - o.l + 1
		return
	}
	mid := (o.l + o.r) >> 1
	if o.left == nil {
		o.left = &CountIntervals{l: o.l, r: mid} // 动态开点
	}
	if o.right == nil {
		o.right = &CountIntervals{l: mid + 1, r: o.r} // 动态开点
	}
	if l <= mid {
		o.left.Add(l, r)
	}
	if mid < r {
		o.right.Add(l, r)
	}
	o.cnt = o.left.cnt + o.right.cnt
}

func (o *CountIntervals) Count() int { return o.cnt }

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */
func main() {

}
