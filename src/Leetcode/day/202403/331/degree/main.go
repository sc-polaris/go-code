package main

/*
	在一棵二叉树中：
	1. 每个空节点 '#'	会提供 0 个出度和 1 个入度
	2. 每个非空节点   会提供 2 个出度和 1 个入度（根节点的入度是 0）
*/

import "strings"

func isValidSerialization(preorder string) bool {
	// 因为，我们加入一个非空节点时，都会对 diff 先减去 1（入度），再加上 2（出度）但是由于根节点没有父节点，所以其入度为 0，出度为 2
	// 因此 diff 初始化为 1，是为了在加入根节点的时候，diff 先减去 1（入度），再加上 2（出度），此时 diff 正好应该是2.
	diff := 1
	for _, s := range strings.Split(preorder, ",") {
		diff--
		if diff < 0 {
			return false
		}
		if s != "#" {
			diff += 2
		}
	}
	return diff == 0
}
