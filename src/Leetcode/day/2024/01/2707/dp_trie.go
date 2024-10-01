package main

type Node struct {
	children [26]*Node
	isEnd    bool
}

func minExtraChar(s string, dictionary []string) int {
	root := &Node{}
	// 逆序
	for _, w := range dictionary {
		node := root
		for k := len(w) - 1; k >= 0; k-- {
			i := w[k] - 'a'
			if node.children[i] == nil {
				node.children[i] = &Node{}
			}
			node = node.children[i]
		}
		node.isEnd = true
	}

	n := len(s)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] + 1 // 不选
		node := root
		for j := i - 1; j >= 0; j-- {
			node = node.children[s[j]-'a']
			if node == nil {
				break
			}
			if node.isEnd {
				f[i] = min(f[i], f[j])
			}
		}
	}

	return f[n]
}

func main() {

}
