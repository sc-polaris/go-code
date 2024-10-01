package main

/*
	给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

	如果子数组中所有元素都相等，则认为子数组是一个 等值子数组 。注意，空数组是 等值子数组 。

	从 nums 中删除最多 k 个元素后，返回可能的最长等值子数组的长度。

	子数组 是数组中一个连续且可能为空的元素序列。
*/

/*
	把相同元素分组，相同元素的下标记录到哈希表（或者数组）posLists 中。
	[1,3,2,3,1,3]
	例如示例 1，元素 3 在 nums 中的下标有 1,3,5，那么 posLists[3] = [1,3,5]。
	遍历 PosLists 中的每个下标列表 pos，例如遍历 pos = [1,3,5]。
	请记住，pos 中保存的是下标，这些下标在 nums 中的对应元素都相同。
	然后用滑动窗口计算。设窗口左右端点为 left 和 right。
	假设 nums 的等值子数组的元素下标从 pos[left] 到 pos[right]，那么在删除前，子数组的长度为为
							pos[right] - pos[left] + 1
	这个子数组有
							right - left + 1
	个数都是相同的，无需删除，其余元素都需要删除，那么需要删除的元素个数就是
							pos[right] - pos[left] - (right - left) =
							pos[right] - right - pos[left] + left =
						 	left - pos[left] - (right - pos[right])
	如果上式大于 k，说明要删除的数太多了，那么移动左指针 left，直到上式小于等于 k，此时用 right - left + 1 更新答案的最大值。

	代码实现时，为简化上式，pos 实际保存的是 pos[i] - i，也就是把上面的每个 pos[i] 都减去其在 pos 中的下标 i，于是需要删除的元素个数简化为
							pos[right] - pos[left]
*/

func longestEqualSubarray(nums []int, k int) (ans int) {
	posLists := make([][]int, len(nums)+1)
	for i, x := range nums {
		posLists[x] = append(posLists[x], i-len(posLists[x]))
	}

	for _, pos := range posLists {
		if len(pos) <= ans { // 无法让 ans 变得更大
			continue
		}
		left := 0
		for right, p := range pos {
			for p-pos[left] > k { // 要删除的数太多了
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}

func main() {
	longestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3)
}
