package main

import (
	"bufio"
	"fmt"
	"os"
)

type TreeNode struct {
	Val         byte
	Left, Right *TreeNode
}

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
)

func newTreeNode(val byte) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

var res []byte

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	dfs(root.Left)
	dfs(root.Right)
}

func main() {
	defer ot.Flush()

	var inInorder, inLorder string

	fmt.Fscan(in, &inInorder, &inLorder)

	inorder := []byte(inInorder)
	lorder := []byte(inLorder)
	pos := make(map[byte]int)
	for i := 0; i < len(inorder); i++ {
		pos[inorder[i]] = i
	}

	st := make([]bool, 26)
	q := make([]*TreeNode, 26)

	q[0] = newTreeNode(lorder[0])        // 根结点
	for i, j := 0, 1; j < len(lorder); { // 按层遍历，i是当前这层的起点，j是下一层的起点
		for end := j; i < end; i++ { // 遍历当前这层
			p := pos[lorder[i]] // 获取当前结点
			st[p] = true        // 标记当前结点使用过
			// 判断左儿子是否存在
			if p-1 >= 0 && !st[p-1] {
				q[i].Left = newTreeNode(lorder[j])
				q[j], j = q[i].Left, j+1
			}
			// 判断右儿子是否存在
			if p+1 < len(lorder) && !st[p+1] {
				q[i].Right = newTreeNode(lorder[j])
				q[j], j = q[i].Right, j+1
			}
		}
	}

	dfs(q[0])

	fmt.Fprintln(ot, string(res))
}
