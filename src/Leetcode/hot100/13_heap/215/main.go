package main

func findKthLargest(nums []int, k int) int {
	var quickSelect func(nums []int, l, r, k int) int
	quickSelect = func(nums []int, l, r, k int) int {
		if l == r {
			return nums[k]
		}
		partition := nums[l]
		i, j := l-1, r+1
		for i < j {
			for i++; nums[i] < partition; i++ {
			}
			for j--; nums[j] > partition; j-- {
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		if k <= j {
			return quickSelect(nums, l, j, k)
		}
		return quickSelect(nums, j+1, r, k)
	}

	n := len(nums)
	return quickSelect(nums, 0, n-1, n-k)
}

func findKthLargest2(nums []int, k int) int {
	buckets := [20001]int{}
	for _, x := range nums {
		buckets[x+10000]++
	}
	for i := 20000; i >= 0; i-- {
		k = k - buckets[i]
		if k <= 0 {
			return i - 10000
		}
	}
	return 0
}
