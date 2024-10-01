package main

import "math/bits"

/*
	根据前缀和的思想，计算从 from 到 to 的幂次之和，等于「前 to+1 个幂次之和」减去「前 from 个幂次之和」。

	计算幂次个数
	定义 ones(n) 为 [0,n−1] 中数字的二进制中的 1 的个数之和。（0 中没有 1，为方便描述，把左边界设为 0。）
	如果 n = 2^i，可以证明：
					ones(2^i) = i*2^(i-1)
	例如 ones(2^2) 为 [0,3] 中数字的二进制中的 1 的个数之和，即 1+1+2=4。
	我们要找一个最大的 n，满足 ones(n)≤k。

	计算幂次之和
	知道了 n，现在来计算幂次之和。
	定义 sumE(n) 为 [0,n−1] 中数字的强数组的幂次之和。（规定 0 的强数组的幂次之和为 0。）
	如果 n=2^i，可以证明：
					sumE(2^i) = i(i-1)/2  * 2^(i-1)

	对于一般的 n，计算方式同「计算幂次个数」，如果 n 二进制从低到高第 i 位是 1，那么幂次之和的增加量，分为如下两部分：
	1. 之前填的 1 的幂次之和 sumI 乘以因为填 1 新增加的元素个数 2^i。
	2. sumE(2^i)。
	可以在「计算幂次个数」的同时计算 sumE。

	得到了幂次之和，可以用快速幂计算 2 的幂模 mod，即为答案。

	代码实现时，乘以 2^(i−1) 可以写成 << (i - 1)。为了避免特判 i=0 的情况，<< (i - 1) 可以用 << i >> 1 代替
*/

func sumE(k int) (res int) {
	var n, cnt1, sumI int
	for i := bits.Len(uint(k+1)) - 1; i >= 0; i-- {
		c := cnt1<<i + i<<i>>1 // 新增的幂次个数
		if c <= k {
			k -= c
			res += sumI<<i + i*(i-1)/2<<i>>1
			sumI += i   // 之前填的 1 的幂次之和
			cnt1++      // 之前填的 1 的个数
			n |= 1 << i // 填1
		}
	}
	// 剩余的 k 个幂次，由 n 的低 k 个 1 补充
	for ; k > 0; k-- {
		res += bits.TrailingZeros(uint(n))
		n &= n - 1 // // 去掉最低位的 1（置为 0）
	}
	return
}

func pow(x, n, mod int) int {
	res := 1 % mod // 注意 mod 可能等于 1
	for ; n > 0; n /= 2 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func findProductsOfElements(queries [][]int64) []int {
	ans := make([]int, len(queries))
	for i, q := range queries {
		er := sumE(int(q[1]) + 1)
		el := sumE(int(q[0]))
		ans[i] = pow(2, er-el, int(q[2]))
	}
	return ans
}
