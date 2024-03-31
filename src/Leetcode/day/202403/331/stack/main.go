package stack

/*
	我们将字符串 preorder 按逗号分割成数组，然后遍历数组，如果遇到了连续两个 '#'，并且第三个元素不是 '#'，
	那么就将这三个元素替换成一个 '#'，这个过程一直持续到数组遍历结束。
	最后，判断数组长度是否为 1，且数组唯一的元素是否为 '#' 即可。
*/

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	var stk []string
	for _, s := range strings.Split(preorder, ",") {
		stk = append(stk, s)
		for len(stk) >= 3 && stk[len(stk)-1] == "#" && stk[len(stk)-2] == "#" && stk[len(stk)-3] != "#" {
			stk = stk[:len(stk)-3]
			stk = append(stk, "#")
		}
	}
	return len(stk) == 1 && stk[0] == "#"
}
