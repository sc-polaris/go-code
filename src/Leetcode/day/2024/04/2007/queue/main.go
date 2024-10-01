package main

import "slices"

func findOriginalArray(changed []int) (ans []int) {
	slices.Sort(changed)
	var q []int
	for _, x := range changed {
		if len(q) > 0 {
			if q[0] < 0 { // 无法配对
				return nil
			}
			if q[0] == x { // 配对成功
				q = q[1:] // 清除一个标记
				continue
			}
		}
		ans = append(ans, x)
		q = append(q, x*2) // 添加双倍标记
	}
	// 只有所有双倍标记都被清除掉，才能说明 changed 是一个双倍数组
	if len(q) > 0 {
		return nil
	}
	return
}
