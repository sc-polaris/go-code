package main

/*
	1. 最终所有元素一定变成了一个在 nums 中的数
	2. 考虑把数字 x 「扩散」到其它位置，那么每一秒 x 都可以向左右扩散一位
		多个相同数字 x 同时扩散，那么扩散完整个数组的耗时，就取决于相距最远的两个相邻的 x。
		假设这两个 x 的下标分别是 i 和 j，且 i < j，那么耗时为：
							(j-i)/2
		枚举不同的 x，计算相应的耗时，更新答案的最小值
	3. 统计所有相同数字的下标，记到一个哈希表 pos 中。
		设 pos[x] 列表第一个下标是 p，最后一个下标是 q。本题数组可以视为换新的，所以 p 和 q 也
		是相邻的，耗时为 [n-(q-p)]/2
		也可以在 pos[x] 列表末尾添加一个 p+n，就可以转换成非环形数组处理了。

*/

func minimumSeconds(nums []int) int {
	pos := make(map[int][]int)
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	n := len(nums)
	ans := n
	for _, a := range pos {
		mx := n - a[len(a)-1] + a[0]
		for i := 1; i < len(a); i++ {
			mx = max(mx, a[i]-a[i-1])
		}
		ans = min(ans, mx)
	}
	return ans / 2
}

func main() {

}
