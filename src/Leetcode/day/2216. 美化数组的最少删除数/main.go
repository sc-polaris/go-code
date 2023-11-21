package _216__美化数组的最少删除数

func minDeletion(nums []int) int {
	// 使用变量 cnt 代表已删除的元素个数，由于每次删除元素，剩余元素都会往前移动，因此当前下标为 i−cnt
	n, cnt := len(nums), 0
	for i := 0; i < n; i++ {
		// 若当前下标为偶数，且与下一位置元素相同，那么当前元素需被删除
		if (i-cnt)%2 == 0 && i+1 < n && nums[i] == nums[i+1] {
			cnt++
		}
	}

	// 最终长度数组n-cnt如果为奇数需要再删除末位元素
	if (n-cnt)%2 != 0 {
		cnt++
	}

	return cnt
}
