package _07

type BinaryIndexTree struct {
	n int
	c []int
}

func NewBinaryIndexTree(n int) *BinaryIndexTree {
	return &BinaryIndexTree{n, make([]int, n+1)}
}

func lowbit(x int) int {
	return x & -x
}

func (t *BinaryIndexTree) update(x, v int) {
	for ; x <= t.n; x += lowbit(x) {
		t.c[x] += v
	}
}

func (t *BinaryIndexTree) query(x int) (s int) {
	for ; x > 0; x -= lowbit(x) {
		s += t.c[x]
	}
	return
}

type NumArray struct {
	tree *BinaryIndexTree
}

func Constructor(nums []int) NumArray {
	tree := NewBinaryIndexTree(len(nums))
	for i, v := range nums {
		tree.update(i+1, v)
	}
	return NumArray{tree: tree}
}

func (t *NumArray) Update(index int, val int) {
	t.tree.update(index+1, val-t.SumRange(index, index))
}

func (t *NumArray) SumRange(left int, right int) int {
	return t.tree.query(right+1) - t.tree.query(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
