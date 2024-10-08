package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (ans []float64) {
	type data struct{ sum, count int }
	var levelData []data
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(levelData) {
			levelData[level].sum += node.Val
			levelData[level].count++
		} else {
			levelData = append(levelData, data{node.Val, 1})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	for _, d := range levelData {
		ans = append(ans, float64(d.sum)/float64(d.count))
	}
	return
}

func averageOfLevels2(root *TreeNode) (ans []float64) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum := 0
		tmp := q
		q = nil
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(len(tmp)))
	}
	return
}
