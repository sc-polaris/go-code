package supper

import (
	"math"
	"strings"
)

func calculate(s string) (ans int) {
	m := map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2, '%': 2, '^': 3}

	nums := []int{0}                   // 为了防止第一个数为负数，先往 nums 加个 0
	var ops []byte                     // 存放所有的操作，包括 +/-
	s = strings.ReplaceAll(s, " ", "") // 将所有的空格去掉
	n := len(s)

	for i := 0; i < n; i++ {
		c := s[i]
		switch c {
		case '(':
			ops = append(ops, c)
		case ')':
			// 计算到最近一个左括号为止
			for len(ops) > 0 {
				if ops[len(ops)-1] != '(' {
					nums, ops = calc(nums, ops)
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		default:
			if '0' <= c && c <= '9' {
				num := 0
				// 将从 i 位置开始后面的连续数字整体取出，加入 nums
				j := i
				for ; j < n && '0' <= s[j] && s[j] <= '9'; j++ {
					num = num*10 + int(s[j]-'0')
				}
				i = j - 1
				nums = append(nums, num)
			} else {
				if i > 0 && (s[i-1] == '(' || s[i-1] == '+' || s[i-1] == '-') {
					nums = append(nums, 0)
				}
				// 有一个新操作要入栈时，先把栈内可以算的都算了
				for len(ops) > 0 && ops[len(ops)-1] != '(' {
					// // 只有满足「栈顶运算符」比「当前运算符」优先级高/同等，才进行运算
					if m[ops[len(ops)-1]] >= m[c] {
						nums, ops = calc(nums, ops)
					} else {
						break
					}
				}
				// 操作符入栈
				ops = append(ops, c)
			}
		}
	}
	// 如果最后一个是数字，读取完之后没有进行计算
	// 此时要把剩余的计算完, 可能有多个都没有计算
	// 例如3+2*2, 所以要用循环
	for len(ops) > 0 {
		nums, ops = calc(nums, ops)
	}
	return nums[len(nums)-1]
}

func calc(nums []int, ops []byte) ([]int, []byte) {
	if len(nums) == 0 || len(nums) < 2 || len(ops) == 0 {
		return nums, ops
	}
	b, a := nums[len(nums)-1], nums[len(nums)-2]
	nums = nums[:len(nums)-2]
	op := ops[len(ops)-1]
	ops = ops[:len(ops)-1]
	ans := 0
	switch op {
	case '+':
		ans = a + b
	case '-':
		ans = a - b
	case '*':
		ans = a * b
	case '/':
		ans = a / b
	case '%':
		ans = a % b
	case '^':
		ans = int(math.Pow(float64(a), float64(b)))
	}
	nums = append(nums, ans)
	return nums, ops
}
