package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
 * 剪枝一：按照从大到小的顺序枚举每一根木棒
 * 剪枝二：木棍内部编号递增
 * 剪枝三：当前木棒失败了，跳过所有相等的木棒
 * 剪枝四：如果在某一根木棍中放第一个木棒失败了，则一定失败
 * 剪枝五：如果在某一根木棍中放最后一个木棒失败了，则一定失败
 */

var (
	in          = bufio.NewReader(os.Stdin)
	ot          = bufio.NewWriter(os.Stdout)
	n           int
	sticks      []int
	st          []bool // 剪枝二：木棍内部编号递增
	sum, length int
)

// 第u组, cur第u组的已有长度, start表示第u组的枚举位置;
func dfs(u, cur, start int) bool {
	if u*length == sum {
		return true
	}
	if cur == length {
		return dfs(u+1, 0, 0)
	}

	for i := start; i < n; i++ {
		l := sticks[i]
		if st[i] || cur+l > length {
			continue
		}

		st[i] = true            // 标记当前木棒已经被使用
		if dfs(u, cur+l, i+1) { // 枚举第i+1根木棒
			return true
		}
		st[i] = false // 还原现场

		// 剪枝四：cur == 0，说明当前拼接的是木棒的第一根木棍，又因为失败了说明是第一根木棍失败了
		// 剪枝五：cur + l == length 说明当前木棍恰好凑满，但是在后面的木棒拼凑时失败了
		if cur == 0 || cur+l == length {
			return false
		}

		// 剪枝三：当前木棒失败了，跳过所有相等的木棒
		j := i
		for j < n && sticks[j] == l {
			j++
		}
		i = j - 1
	}

	return false
}

func main() {
	defer ot.Flush()
	for {
		fmt.Fscan(in, &n)

		if n == 0 {
			break
		}

		sticks = make([]int, n)
		sum, length = 0, 0
		for i := range sticks {
			var l int
			fmt.Fscan(in, &l)
			sticks[i] = l
			sum += l
			if length < l {
				length = l
			}
		}

		// 剪枝一：按照从大到小的顺序枚举每一根木棒
		sort.Slice(sticks, func(i, j int) bool {
			return sticks[i] > sticks[j]
		})

		st = make([]bool, n)
		for {
			if sum%length == 0 && dfs(0, 0, 0) {
				fmt.Fprintln(ot, length)
				break
			}
			length++
		}
	}
}
