package main

// 方法一：前后缀分解
func trap(height []int) (ans int) {
	n := len(height)
	preMax := make([]int, n) // preMax[i] 表示从 height[0] 到 height[i] 的最大值
	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}

	sufMax := make([]int, n) // sufMax[i] 表示从 height[i] 到 height[n-1] 的最大值
	sufMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(preMax[i], sufMax[i]) - h // 累加每个水桶能接多少水
	}
	return
}

// 方法二：相向双指针
func trap2(height []int) (ans int) {
	left, right, preMax, sufMax := 0, len(height)-1, 0, 0
	// 注意 while 循环可以不加等号，因为在「谁小移动谁」的规则下，相遇的位置一定是最高的柱子，这个柱子是无法接水的。
	for left < right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
	}
	return
}

// 方法三：单调栈 相当于「横着」计算面积。
// 这个方法可以总结成 16 个字：找上一个更大元素，在找的过程中填坑。
// 注意 while 中加了等号，这可以让栈中没有重复元素，从而在有很多重复元素的情况下，使用更少的空间。
func trap3(height []int) (ans int) {
	var st []int
	for i, h := range height {
		for len(st) > 0 && h >= height[st[len(st)-1]] {
			bottomH := height[st[len(st)-1]]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			dh := min(height[left], h) - bottomH // 面积的高
			ans += dh * (i - left - 1)
		}
		st = append(st, i)
	}
	return
}
