package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func quickSelect(q []int, l, r, k int) int {
	if l >= r {
		return q[k]
	}

	i, j, x := l-1, r+1, q[(l+r)>>1]
	for i < j {
		for i = i + 1; q[i] < x; i++ {
		}
		for j = j - 1; q[j] > x; j-- {
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	if j >= k {
		return quickSelect(q, l, j, k)
	} else {
		return quickSelect(q, j+1, r, k)
	}
}

func getMidNum(q []int) int {
	n := len(q)
	if n&1 == 1 { // 奇数
		return quickSelect(q, 0, n-1, n>>1)
	} else {
		return (quickSelect(q, 0, n-1, n>>1) + quickSelect(q, 0, n-1, (n-1)>>1)) >> 1
	}
}

func main() {
	defer ot.Flush()

	var n int
	fmt.Fscan(in, &n)

	// 预处理前缀和
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		s[i] += s[i-1]
	}

	avg := s[n] / n // 平均数
	var c []int
	for i := 1; i < n; i++ {
		c = append(c, avg*i-s[i])
	}
	c = append(c, 0) // c[n] = 0;

	//sort.Ints(c)

	//md := c[len(c)>>1] // 求中位数
	md := getMidNum(c) // 求中位数 快速选择优化
	var res int
	for _, v := range c {
		res += abs(v - md)
	}

	fmt.Fprintln(ot, res)
}
