package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type Node struct {
	Time   int
	Profit int
}

const N = 10010

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
	p  [N]int
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 并查集：路径压缩, 把这个集合中所有的 都指向根
func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}

	return p[x]
}

func main() {
	defer ot.Flush()

	for {
		_, err := fmt.Fscan(in, &n)
		if err == io.EOF {
			break
		}
		var d int
		products := make([]Node, n)
		for i := range products {
			fmt.Fscan(in, &products[i].Profit, &products[i].Time)
			d = max(d, products[i].Time) // 取得一个最晚过期时间
		}

		// 按照利润从大到小排序
		sort.Slice(products, func(i, j int) bool {
			return products[i].Profit > products[j].Profit
		})

		for i := 0; i <= d; i++ { // 起初每一天各自构成一个集合
			p[i] = i
		}

		var res int
		// 利用路径压缩，可以快速找出从过期时间往前数第一个空闲的天数
		for i := 0; i < n; i++ { // 从利润大的开始
			r := find(products[i].Time) // 获取利润最大的商品的过期日期
			if r > 0 {                  // 如果这个"位置"还没有被用掉
				res += products[i].Profit // 更新答案
				p[r] = r - 1              // 合并两个集合(r与r-1)-> 把这个"位置"指向他前一个"位置"
			}
		}

		fmt.Fprintln(ot, res)
	}
}
