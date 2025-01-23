package main

type Tire struct {
	son  [26]*Tire
	word string
}

func (t *Tire) Insert(word string) {
	node := t
	for _, c := range word {
		c -= 'a'
		for node.son[c] == nil {
			node.son[c] = new(Tire)
		}
		node = node.son[c]
	}
	node.word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func findWords(board [][]byte, words []string) []string {
	t := &Tire{}
	for _, word := range words {
		t.Insert(word)
	}

	m, n := len(board), len(board[0])
	vis := make(map[string]bool)

	var dfs func(*Tire, int, int)
	dfs = func(node *Tire, x, y int) {
		c := board[x][y]
		node = node.son[c-'a']
		if node == nil {
			return
		}
		if node.word != "" {
			vis[node.word] = true
		}
		board[x][y] = '#'
		for _, d := range dirs {
			dx, dy := x+d.x, y+d.y
			if 0 <= dx && dx < m && 0 <= dy && dy < n && board[dx][dy] != '#' {
				dfs(node, dx, dy)
			}
		}
		board[x][y] = c // 回溯
	}
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}
	var ans []string
	for s := range vis {
		ans = append(ans, s)
	}
	return ans
}
