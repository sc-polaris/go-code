package main

type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.q2 = append(s.q2, x)
	for len(s.q1) > 0 {
		s.q2 = append(s.q2, s.q1[0])
		s.q1 = s.q1[1:]
	}
	s.q1, s.q2 = s.q2, s.q1
}

func (s *MyStack) Pop() int {
	x := s.q1[0]
	s.q1 = s.q1[1:]
	return x
}

func (s *MyStack) Top() int {
	return s.q1[0]
}

func (s *MyStack) Empty() bool {
	return len(s.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {

}
