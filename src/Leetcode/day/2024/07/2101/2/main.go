package main

import "math/bits"

/*
	方法二：Bitset 优化 Floyd

	Floyd 算法可以计算任意两点间的最短路长度。但如果只要求计算任意两点间「能否」到达，则可以用 bool 数组代替 int 数组。
	进一步地，可以把 bool 数组压缩成 bitset。

	定义 f[i] 表示 i 可以到达的节点集合。注意 i 一定在 f[i] 中。
	如果 i 可以到达 k，那么 k 能到达的点，i 也可以到达，于是 i 能到达的点就是 f[i] 和 f[k] 的并集，即
							f[i] = f[i] ∪ f[k]
*/

func maximumDetonation(bombs [][]int) (ans int) {
	n := len(bombs)
	f := make([]bitset, n)
	for i, p := range bombs {
		x, y, r := p[0], p[1], p[2]
		f[i] = newBitset(n)
		for j, q := range bombs {
			dx := x - q[0]
			dy := y - q[1]
			if dx*dx+dy*dy <= r*r {
				f[i].set(j) // i 可以到达 j
			}
		}
	}

	for k, fk := range f {
		for _, fi := range f {
			if fi.has(k) { // i 可以到达 k
				fi.or(fk) // i 也可以到 k 可以到达的点
			}
		}
	}

	for _, s := range f {
		ans = max(ans, s.onesCount()) // 集合大小的最大值
	}
	return
}

const w = bits.UintSize

type bitset []uint

func newBitset(n int) bitset {
	return make(bitset, (n+w-1)/w) // 需要 ceil(n/w) 个 w 位整数
}

func (b bitset) set(p int) {
	b[p/w] |= 1 << (p % w)
}

func (b bitset) has(p int) bool {
	return b[p/w]&(1<<(p%w)) != 0
}

func (b bitset) or(other bitset) {
	for i, x := range other {
		b[i] |= x
	}
}

func (b bitset) onesCount() (c int) {
	for _, x := range b {
		c += bits.OnesCount(x)
	}
	return
}
