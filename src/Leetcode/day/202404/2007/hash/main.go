package hash

import "slices"

func findOriginalArray(changed []int) (ans []int) {
	slices.Sort(changed)
	cnt := make(map[int]int)
	for _, x := range changed {
		if cnt[x] == 0 { // x 不是双倍后的元素
			cnt[x*2]++ // 标记一个双倍元素
			ans = append(ans, x)
		} else { // x 是双倍后的元素
			cnt[x]-- // 清除一个标记
			if cnt[x] == 0 {
				delete(cnt, x)
			}
		}
	}
	// 只有所有双倍标记都被清除掉，才能说明 changed 是一个双倍数组
	if len(cnt) > 0 {
		return nil
	}
	return
}
