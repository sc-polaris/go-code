package main

import "container/list"

/*
	在考场里，一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1 。

	当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。
	如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 0 号座位上。)

	返回 ExamRoom(int N) 类，它有两个公开的函数：其中，
	函数 ExamRoom.seat() 会返回一个 int （整型数据），代表学生坐的位置；
	函数 ExamRoom.leave(int p) 代表坐在座位 p 上的学生现在离开了考场。
	每次调用 ExamRoom.leave(p) 时都保证有学生坐在座位 p 上。
*/

/*
	链表版
	插入和删除都是 o(n)

	首先需要考虑的点有

	新的位置需要放在两个已知点之间，如何求出这个点。因为要距离最大化，那么肯定是在中点，如果距离是奇数，那就是中点，如果是偶数，那就直接放在左侧。
	如果要插入到头或者尾，需要特殊判断距离
	在找到最大距离的同时，需要将对应的插入位置也保存下来
	删除的时候只要遍历链表删除即可
*/

type ExamRoom struct {
	seat *list.List
	n    int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{seat: list.New(), n: n - 1}
}

func (er *ExamRoom) Seat() int {
	// 还没有人入座
	if er.seat.Len() == 0 {
		er.seat.PushFront(0)
		return 0
	}
	e := er.seat.Front()
	pre := e.Value.(int)
	mx := pre // 头部特殊判断
	addVal := 0
	addFront := e
	e = e.Next()
	for ; e != nil; e = e.Next() {
		val := e.Value.(int)
		distance := (val - pre) / 2 // 两点之间的最远距离
		if distance > mx {
			mx = distance
			addFront = e // 需要插入的点的后一个元素。方便找到后直接插入
			addVal = pre + distance
		}
		pre = val
	}
	distance := er.n - pre // 尾部特殊判断
	if distance > mx {
		er.seat.PushBack(er.n) // 直接插入到链表尾部
		return er.n
	}
	er.seat.InsertBefore(addVal, addFront) // 插入
	return addVal
}

func (er *ExamRoom) Leave(p int) {
	for e := er.seat.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			er.seat.Remove(e)
			return
		}
	}
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
