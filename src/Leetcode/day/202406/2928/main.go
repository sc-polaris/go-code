package main

/*
	给你两个正整数 n 和 limit 。

	请你将 n 颗糖果分给 3 位小朋友，确保没有任何小朋友得到超过 limit 颗糖果，请你返回满足此条件下的 总方案数 。
*/

/*
	所有方案数 - 至少一个小朋友分到的糖果超过 limit - 至少两个小朋友分到的糖果超过 limit - 三个小朋友分到的糖果超过 limit

	所有方案数：相当于把 n 个无区别的小球放入 3 个有区别的盒子，允许空盒的方案数。
	隔板法：假设 n 个球和 2 个隔板放到 n+2 个位置，第一个隔板前的球放入第一个盒子，第一个隔板和第二个隔板之间的球放入第二个
		   盒子，第二个隔板后的球放入第三个盒子。那么从 n+2 个位置中选 2 个位置放隔板，有 C(n+2,2) 种放法。
	注意隔板可以放在最左边或最右边，也可以连续放，对应着空盒的情况。例如第一个隔板放在最左边，意味着第一个盒子是空的；又例如第
	一个隔板和第二个隔板相邻，意味着第二个盒子是空的。

	至少一个小朋友分到的糖果超过 limit：
	设三个小朋友分别叫 A，B，C
	只关注 A。如果 A 分到的糖果超过 limit，那么先分给他 limit + 1 颗糖果，问题剩下 n-(limit+1) 颗糖果分给三个小朋友的方案数，
	即 C(n-(limit+1)+2,2)。注意 B 和 C 分到的糖果是否超过 limit 我们是不关注的。
	只关注 B 和只关注 C 同上。加起来就是 3*C(n-(limit+1)+2,2),但这样就重复统计了「至少两个小朋友分到的糖果超过 limit」的情况，要减去。

	至少两个小朋友分到的糖果超过 limit：
	只关注 A，B 方案数为 C(n-2*(limit+1)+2,2) = C(n-2*limit,2)
	直接加起来，就是 3*C(n-2*(limit+1)+2,2) 但这样就重复统计了「三个小朋友分到的糖果均超过 limit」的情况，要减去。

	三个小朋友分到的糖果超过 limit：
	C(n-3(limit+1)+2,2)

	总结：
	不合法的方案数为「至少一个」减去「至少两个」加上「三个」，这就是容斥原理。
	C(n+2,2) - (3*C(n-(limit+1)+2,2 - (3*C(n-2*limit,2) - C(n-3(limit+1)+2,2)))  =
	C(n+2,2) - (3*C(n-(limit+1)+2,2 - 3*C(n-2*limit,2) + C(n-3(limit+1)+2,2))	  =
	C(n+2,2) - 3*C(n-(limit+1)+2,2 + 3*C(n-2*limit,2) - C(n-3(limit+1)+2,2))

*/

func c2(n int) int {
	if n < 2 {
		return 0
	}
	return n * (n - 1) / 2
}

func distributeCandies(n int, limit int) int {
	return c2(n+2) - 3*c2(n-limit+1) + 3*c2(n-2*limit) - c2(n-3*limit-1)
}
