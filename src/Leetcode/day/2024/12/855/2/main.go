package main

import "github.com/emirpasic/gods/trees/redblacktree"

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
	考虑到每次 seat() 时都需要找到最大距离的座位，我们可以使用有序集合来保存座位区间。有序集合的每个元素为一个二元组 (l,r)，
	表示 l 和 r 之间（不包括 l 和 r）的座位可以坐学生。初始时有序集合中只有一个元素 (−1,n)，表示 (−1,n) 之间的座位可以坐学生。

	另外，我们使用两个哈希表 left 和 right 来维护每个有学生的座位的左右邻居学生，方便我们在 leave(p) 时合并两个座位区间。
*/

type pair struct{ x, y int }

type ExamRoom struct {
	t     *redblacktree.Tree
	left  map[int]int
	right map[int]int
	n     int
}

func Constructor(n int) ExamRoom {
	dist := func(p pair) int {
		l, r := p.x, p.y
		if l == -1 || r == n {
			return r - l - 1
		}
		return (r - l) >> 1
	}
	cmp := func(a, b any) int {
		x, y := a.(pair), b.(pair)
		d1, d2 := dist(x), dist(y)
		if d1 == d2 {
			return x.x - y.x
		}
		return d2 - d1
	}
	er := ExamRoom{redblacktree.NewWith(cmp), make(map[int]int), make(map[int]int), n}
	er.add(pair{-1, n})
	return er
}

func (er *ExamRoom) Seat() int {
	s := er.t.Left().Key.(pair)
	p := (s.x + s.y) >> 1
	if s.x == -1 {
		p = 0
	} else if s.y == er.n {
		p = er.n - 1
	}
	er.del(s)
	er.add(pair{s.x, p})
	er.add(pair{p, s.y})
	return p
}

func (er *ExamRoom) Leave(p int) {
	l, _ := er.left[p]
	r, _ := er.right[p]
	er.del(pair{l, p})
	er.del(pair{p, r})
	er.add(pair{l, r})
}

func (er *ExamRoom) add(s pair) {
	er.t.Put(s, struct{}{})
	er.left[s.y] = s.x
	er.right[s.x] = s.y
}

func (er *ExamRoom) del(s pair) {
	er.t.Remove(s)
	delete(er.left, s.y)
	delete(er.right, s.x)
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
