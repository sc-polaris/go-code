package main

/*
	给你一个下标从 0 开始的整数数组 nums 。如果下标对 i、j 满足 0 ≤ i < j < nums.length ，如果 nums[i] 的
	第一个数字 和 nums[j] 的 最后一个数字 互质 ，则认为 nums[i] 和 nums[j] 是一组 美丽下标对 。

	返回 nums 中 美丽下标对 的总数目。

	对于两个整数 x 和 y ，如果不存在大于 1 的整数可以整除它们，则认为 x 和 y 互质 。换而言之，如果 gcd(x, y) == 1 ，
	则认为 x 和 y 互质，其中 gcd(x, y) 是 x 和 y 的 最大公因数 。
*/

/*
	由于 nums[i] 的最高位在 [1,9] 中，我们可以在遍历数组的同时，统计最高位出现的次数，这样就只需枚举 [1,9] 中的与
	x mod 10 互质的数，把对应的出现次数加到答案中。

	具体算法如下：
	1. 初始化答案 ans = 0，初始化长为 10 的 cnt 数组，初始值均为 0.
	2. 遍历 nums，设 x = nums[j]。
	3. 枚举 [1,9] 内的数字 y，如果与 x mod 10 互质，则 ans 增加 cnt[y]。
	4. 计算 x 的最高位，将其出现次数加 一。
	5. 返回 ans
*/

func countBeautifulPairs(nums []int) (ans int) {
	cnt := [10]int{}
	for _, x := range nums {
		for y := 1; y < 10; y++ {
			if cnt[y] > 0 && gcd(x%10, y) == 1 {
				ans += cnt[y]
			}
		}
		for x >= 10 { // 这里需要 O(log x) 的时间
			x /= 10
		}
		cnt[x]++ // 统计最高位出现的次数
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
