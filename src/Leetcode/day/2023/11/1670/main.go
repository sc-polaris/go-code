package main

import "container/list"

type FrontMiddleBackQueue struct {
	left  *list.List
	right *list.List
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{left: list.New(), right: list.New()}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {
	q.left.PushFront(val)
	if q.left.Len() == q.right.Len()+2 {
		q.right.PushFront(q.left.Back().Value.(int))
		q.left.Remove(q.left.Back())
	}
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	if q.left.Len() == q.right.Len()+1 {
		q.right.PushFront(q.left.Back().Value.(int))
		q.left.Remove(q.left.Back())
	}
	q.left.PushBack(val)
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	q.right.PushBack(val)
	if q.left.Len()+1 == q.right.Len() {
		q.left.PushBack(q.right.Front().Value.(int))
		q.right.Remove(q.right.Front())
	}
}

func (q *FrontMiddleBackQueue) PopFront() int {
	if q.left.Len() == 0 {
		return -1
	}
	val := q.left.Front().Value.(int)
	q.left.Remove(q.left.Front())
	if q.left.Len()+1 == q.right.Len() {
		q.left.PushBack(q.right.Front().Value.(int))
		q.right.Remove(q.right.Front())
	}
	return val
}

func (q *FrontMiddleBackQueue) PopMiddle() int {
	if q.left.Len() == 0 {
		return -1
	}
	val := q.left.Back().Value.(int)
	q.left.Remove(q.left.Back())
	if q.left.Len()+1 == q.right.Len() {
		q.left.PushBack(q.right.Front().Value.(int))
		q.right.Remove(q.right.Front())
	}
	return val
}

func (q *FrontMiddleBackQueue) PopBack() int {
	if q.left.Len() == 0 {
		return -1
	}
	if q.right.Len() == 0 {
		val := q.left.Back().Value.(int)
		q.left.Remove(q.left.Back())
		return val
	} else {
		val := q.right.Back().Value.(int)
		q.right.Remove(q.right.Back())
		if q.left.Len() == q.right.Len()+2 {
			q.right.PushFront(q.left.Back().Value.(int))
			q.left.Remove(q.left.Back())
		}
		return val
	}
}

/**
* Your FrontMiddleBackQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.PushFront(val);
* obj.PushMiddle(val);
* obj.PushBack(val);
* param_4 := obj.PopFront();
* param_5 := obj.PopMiddle();
* param_6 := obj.PopBack();
 */

func main() {

}
