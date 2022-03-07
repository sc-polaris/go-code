package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// 划分算法
// 返回基数划分后所在的下标
func partition(nums []int, left, right, pivotIndex int) int {
	pivot := nums[pivotIndex]
	swap(nums, pivotIndex, right) // 将基数放置最后
	save := left

	for i := left; i < right; i++ {
		if nums[i] < pivot {
			swap(nums, i, save)
			save++
		}
	}

	// 将基数和save所指的数据交换
	swap(nums, save, right)

	return save
}

// 快速选择算法
func quickSelect(nums []int, left, right, k int) int {
	if left == right { // 只有一个元素时
		return nums[left] // 直接返回
	}

	pivotIndex := (left + right) >> 1

	// 进行划分
	pivotIndex = partition(nums, left, right, pivotIndex)

	if pivotIndex == k {
		return nums[k]
	} else if pivotIndex < k {
		return quickSelect(nums, pivotIndex+1, right, k)
	} else {
		return quickSelect(nums, left, pivotIndex-1, k)
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var mid int
	// 奇数直接取中位数
	if n&1 == 1 {
		mid = quickSelect(a, 0, n-1, (n-1)>>1)
	} else { // 偶数取中间两个数的和
		mid = (quickSelect(a, 0, n-1, n>>1) + quickSelect(a, 0, n-1, (n-1)>>1)) >> 1
	}

	var res int
	for i := range a {
		res += abs(a[i] - mid)
	}

	fmt.Fprintln(ot, res)
}
