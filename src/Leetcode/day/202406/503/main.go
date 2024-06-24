package main

/*
	给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。
	数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。
	如果不存在，则输出 -1 。
*/

/*
	本题 nums 是一个循环数组，nums[n-1] 右边是 nums[0]。我们可以把 nums 复制一份，拼在 nums 右边，这样就把环形数组变成
	一般数组了。例如 [1,2,1] 变成 [1,2,1,1,2,1]。

	从右到左

	从右往左遍历，栈中记录下一个更大元素的「候选项」。
	由于左边更大元素会「挡住」右边更小的元素，所以右边更小的元素是无用信息（不会成为左边元素的下一个更大元素），这会导致栈底
	（右边）大，栈顶（左边）小。

	代码实现时，无需真的把数组复制一份，而是用下标模 n 的方式取到对应的元素值。
*/

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var st []int
	for i := 2*n - 1; i >= 0; i-- {
		x := nums[i%n]
		for len(st) > 0 && x >= st[len(st)-1] {
			// 由于 x 的出现，栈顶元素永远不会是左边元素的「下一个更大元素」
			st = st[:len(st)-1]
		}
		if i < n && len(st) > 0 {
			ans[i] = st[len(st)-1]
		}
		st = append(st, x)
	}
	return ans
}

/*
	从左到右
	栈中记录还没算出「下一个更大元素」的那些数的下标。
	只要遍历到比栈顶元素值更大的数，就意味着栈顶元素找到了答案，记录答案，然后从栈顶弹出。
*/

func nextGreaterElements2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var st []int
	for i := 0; i < 2*n; i++ {
		x := nums[i%n]
		for len(st) > 0 && x > nums[st[len(st)-1]] {
			// x 是 nums[st[len(st)-1]] 的下一个更大元素
			// 既然 nums[st[len(st)-1]] 已经算出答案，则从栈顶弹出
			ans[st[len(st)-1]] = x
			st = st[:len(st)-1]
		}
		if i < n {
			st = append(st, i)
		}
	}
	return ans
}
