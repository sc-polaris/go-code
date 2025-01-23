package main

type Node struct {
	son [26]*Node
	end bool
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{&Node{}}
}

func (wd *WordDictionary) AddWord(word string) {
	cur := wd.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil {
			cur.son[c] = new(Node)
		}
		cur = cur.son[c]
	}
	cur.end = true
}

func (wd *WordDictionary) Search(word string) bool {
	var dfs func(int, *Node) bool
	dfs = func(index int, node *Node) bool {
		if index == len(word) {
			return node.end
		}
		c := word[index]
		if c != '.' {
			son := node.son[c-'a']
			if son != nil && dfs(index+1, son) {
				return true
			}
		} else {
			for _, son := range node.son {
				if son != nil && dfs(index+1, son) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, wd.root)
}

func (wd *WordDictionary) SearchBfs(word string) bool {
	type bfsNode struct {
		node  *Node
		index int
	}
	q := []bfsNode{{wd.root, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		node, index := p.node, p.index
		if index == len(word) {
			if node.end {
				return true
			}
			continue
		}

		c := word[index]
		if c != '.' {
			son := node.son[c-'a']
			if son != nil {
				q = append(q, bfsNode{son, index + 1})
			}
		} else {
			for _, son := range node.son {
				if son != nil {
					q = append(q, bfsNode{son, index + 1})
				}
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
