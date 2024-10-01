package main

func minIncrements(n int, cost []int) (ans int) {
	for i := n / 2; i > 0; i-- { // 从最后一个非叶节点开始算
		l, r := cost[i*2-1], cost[i*2]
		if l > r { // 保证 l <= r
			l, r = r, l
		}
		ans += r - l   // 两个子节点变成一样的
		cost[i-1] += r // 累加路径和
	}
	return
}

func main() {}
