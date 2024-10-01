package main

/*
	在 X轴 上有一些奖品。给你一个整数数组 prizePositions ，它按照 非递减 顺序排列，其中 prizePositions[i] 是第 i 件
	奖品的位置。数轴上一个位置可能会有多件奖品。再给你一个整数 k 。

	你可以同时选择两个端点为整数的线段。每个线段的长度都必须是 k 。你可以获得位置在任一线段上的所有奖品（包括线段的两个端点）。
	注意，两个线段可能会有相交。

	比方说 k = 2 ，你可以选择线段 [1, 3] 和 [2, 4] ，你可以获得满足 1 <= prizePositions[i] <= 3 或者
	2 <= prizePositions[i] <= 4 的所有奖品 i 。

	请你返回在选择两个最优线段的前提下，可以获得的 最多 奖品数目。
*/

/*
	方法一：枚举右 维护左
	一条线段：
	从特殊到一般，先想想只有一条线段要怎么做

	如果线段的右端点没有奖品，我们可以把线段左移，使其右端点恰好有奖品，这不会让线段覆盖的奖品个数变少。所以只需枚举
	prizePositions[right] 为线段的右端点，然后需要算出最远（最小）覆盖的奖品的位置 prizePositions[left]，
	此时覆盖的奖品的个数为
								right−left+1
	由于 right 变大时，left 也会变大，有单调性，可以用滑动窗口快速算出 left。
	⚠注意：prizePositions[left] 不一定是线段的左端点。prizePositions[left] 只是最左边的被线段覆盖的那个奖品的位置，
		  线段左端点可能比 prizePositions[left] 更小。

	两条线段：
	两条线段一左一右。考虑枚举右（第二条线段），同时维护左（第一条线段）能覆盖的最多奖品个数。
	贪心地想，两条线段不相交肯定比相交更好，覆盖的奖品可能更多。
	设第二条线段右端点在 prizePositions[right] 时，最远（最小）覆盖的奖品的位置为 prizePositions[left]。
	我们需要计算在 prizePositions[left] 左侧的第一条线段最多可以覆盖多少个奖品。这可以保证两条线段不相交。
	定义 mx[i+1] 表示第一条线段右端点 ≤ prizePositions[i] 时，最多可以覆盖多少个奖品。特别地，定义 mx[0]=0。

	根据 mx 的定义，我们相当于在计算 i−lefti+1 的前缀最大值，其中 lefti 表示右端点覆盖奖品 prizePositions[i] 时，最左边的被线段覆盖的奖品。
	所以有
								mx[i+1] = max(mx[i],i-lefti+1
	如何计算两条线段可以覆盖的奖品个数？
	1. 第二条线段覆盖的奖品个数为 right−left+1。
	2. 第一条线段覆盖的奖品个数为线段右端点 ≤prizePositions[left−1] 时，最多覆盖的奖品个数，即 mx[left]。
	综上，两条线段可以覆盖的奖品个数为
								mx[left]+right-left+1\

	我们遍历了所有的奖品作为第二条线段的右端点，通过 mx[left] 保证第一条线段与第二条线段不相交，且第一条线段覆盖了第二条线段
	 左侧的最多奖品。那么这样遍历后，算出的答案就一定是所有情况中的最大值。

	⚠注意：可以在计算第二条线段的滑动窗口的同时，更新和第一条线段有关的 mx。这是因为两条线段一样长，第二条线段移动到 right
		  时所覆盖的奖品个数，也是第一条线段移动到 right 时所覆盖的奖品个数。

	小优化：如果 2k+1≥prizePositions[n−1]−prizePositions[0]，说明所有奖品都可以被覆盖，直接返回 n。
*/

func maximizeWin(prizePositions []int, k int) (ans int) {
	n := len(prizePositions)
	if 2*k+1 >= prizePositions[n-1]-prizePositions[0] {
		return n
	}
	mx := make([]int, n+1)
	l := 0
	for r, p := range prizePositions {
		for p-prizePositions[l] > k {
			l++
		}
		ans = max(ans, mx[l]+r-l+1)
		mx[r+1] = max(mx[r], r-l+1)
	}
	return
}

/*
	方法二：换一个角度
	两条线段一共涉及到 4 个下标：
	1. 第一条线段覆盖的最小奖品下标。
	2. 第一条线段覆盖的最大奖品下标。
	3. 第二条线段覆盖的最小奖品下标。
	4. 第二条线段覆盖的最大奖品下标。

	考虑「枚举中间」，也就是第一条线段覆盖的最大奖品下标，和第二条线段覆盖的最小奖品下标。

	第一条线段
	写一个和方法一一样的滑动窗口：
	· 枚举覆盖的最大奖品下标为 right，维护覆盖的最小奖品下标 left。
	· 向右移动 right，如果发现 prizePositions[right]−prizePositions[left]>k，就向右移动 left。
	· 循环结束时，覆盖的奖品个数为 right−left+1。

	第二条线段
	仍然是滑动窗口，但改成枚举 left，维护 right。
	· 向右移动 left，如果发现 prizePositions[right]−prizePositions[left]≤k，就向右移动 right。
	· 循环结束时，right−1 是覆盖的最大奖品下标，覆盖的奖品个数为 right−left。

	合二为一
	枚举 mid，既作为第一条线段的 right，又作为第二条线段的 left。
	同方法一，用滑动窗口枚举第二条线段，同时维护第一条线段能覆盖的最多奖品个数 mx。
	枚举 mid：
	1. 首先，跑第二条线段的滑动窗口。
	2. 用 mx+right−mid 更新答案的最大值。
	3. 然后，跑第一条线段的滑动窗口。
	4. 用 mid−left+1 更新 mx 的最大值。
	⚠注意：不能先跑第一条线段的滑动窗口，否则 mx+right−mid 可能会把 mid 处的奖品计入两次。
*/

func maximizeWin2(prizePositions []int, k int) (ans int) {
	n := len(prizePositions)
	if 2*k+1 >= prizePositions[n-1]-prizePositions[0] {
		return n
	}
	mx, l, r := 0, 0, 0
	for mid, p := range prizePositions {
		// 把 prizePositions[mid] 视作第二条线段的左端点，计算第二条线段可以覆盖的最大奖品下标
		for r < n && prizePositions[r]-p <= k {
			r++
		}
		// 循环结束后，right-1 是第二条线段可以覆盖的最大奖品下标
		ans = max(ans, mx+r-mid)
		// 把 prizePositions[mid] 视作第一条线段的右端点，计算第一条线段可以覆盖的最小奖品下标
		for p-prizePositions[l] > k {
			l++
		}
		// 循环结束后，left 是第一条线段可以覆盖的最小奖品下标
		mx = max(mx, mid-l+1)
	}
	return
}
