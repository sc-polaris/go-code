package main

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	var st1, st2 []int
	for i := 0; i < n; i++ {
		v := nums[i]
		// st2 非空且栈顶元素小于当前遍历的元素时，说明当前元素为栈顶元素的「第二大」的整数
		for len(st2) > 0 && nums[st2[len(st2)-1]] < v {
			res[st2[len(st2)-1]] = v
			st2 = st2[:len(st2)-1]
		}
		pos := len(st1) - 1
		for pos >= 0 && nums[st1[pos]] < v {
			pos--
		}
		st2 = append(st2, st1[pos+1:]...)
		st1 = append(st1[:pos+1], i)
	}
	return res
}

func main() {

}
