package main

import "sort"

/*
	实现支持下列接口的「快照数组」- SnapshotArray：

	SnapshotArray(int length) - 初始化一个与指定长度相等的 类数组 的数据结构。初始时，每个元素都等于 0。
	void set(index, val) - 会将指定索引 index 处的元素设置为 val。
	int snap() - 获取该数组的快照，并返回快照的编号 snap_id（快照号是调用 snap() 的总次数减去 1）。
	int get(index, snap_id) - 根据指定的 snap_id 选择快照，并返回该快照指定索引 index 的值。
*/

type pair struct{ snapId, val int }

type SnapshotArray struct {
	curSnapId int
	history   map[int][]pair // 每个 index 的历史修改记录
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{history: make(map[int][]pair)}
}

func (sa *SnapshotArray) Set(index int, val int) {
	sa.history[index] = append(sa.history[index], pair{snapId: sa.curSnapId, val: val})
}

func (sa *SnapshotArray) Snap() int {
	sa.curSnapId++
	return sa.curSnapId - 1
}

func (sa *SnapshotArray) Get(index int, snapId int) int {
	h := sa.history[index]
	// 找快照编号 <= snapId 的最后一次修改记录
	// 等价于找快照编号 >= snapId+1 的第一个修改记录，它的上一个就是答案
	j := sort.Search(len(h), func(j int) bool { return h[j].snapId >= snapId+1 }) - 1
	if j >= 0 {
		return h[j].val
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
