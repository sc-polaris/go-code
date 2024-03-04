package main

type MyQueue struct {
	inStk, outStk []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStk = append(q.inStk, x)
}

func (q *MyQueue) in2Out() {
	for len(q.inStk) > 0 {
		q.outStk = append(q.outStk, q.inStk[len(q.inStk)-1])
		q.inStk = q.inStk[:len(q.inStk)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStk) == 0 {
		q.in2Out()
	}
	x := q.outStk[len(q.outStk)-1]
	q.outStk = q.outStk[:len(q.outStk)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStk) == 0 {
		q.in2Out()
	}
	return q.outStk[len(q.outStk)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStk) == 0 && len(q.outStk) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {

}
