package main

import (
	"math/bits"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	return FindElements{root}
}

func (f *FindElements) Find(target int) bool {
	target++
	cur := f.root                                      // 从根节点出发
	for i := bits.Len(uint(target)) - 2; i >= 0; i-- { // 从次高位开始枚举
		if target>>i&1 == 0 {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
		if cur == nil { // 走到空节点，说明 target 不在二叉树中
			return false
		}
	}
	return true // 没有走到空节点，说明 target 在二叉树中
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

func main() {}
