package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
)

func bfs(start string) int {
	end := "12345678x"

	if start == end {
		return 0
	}

	d := make(map[string]int)
	q := make([]string, 0)
	q = append(q, start)
	d[start] = 0

	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}

	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		k := strings.Index(t, "x")
		x, y := k/3, k%3 // 一维坐标转化二维坐标

		for i := 0; i < 4; i++ {
			temp := []byte(t)
			a, b := x+dx[i], y+dy[i]
			if 0 <= a && a < 3 && 0 <= b && b < 3 {
				temp[a*3+b], temp[x*3+y] = temp[x*3+y], temp[a*3+b]
				state := string(temp)
				if _, ok := d[state]; !ok {
					d[state] = d[t] + 1
					if state == end {
						return d[state]
					}
					q = append(q, state)
				}
			}
		}
	}

	return -1
}

func main() {
	defer ot.Flush()
	var start, state string
	var c string
	for len(start) < 9 {
		fmt.Fscan(in, &c)
		start += c
		if c != "x" {
			state += c
		}
	}

	cnt := 0
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if state[i] > state[j] { // 逆序的数量
				cnt++
			}
		}
	}

	if cnt%2 == 1 {
		fmt.Fprintln(ot, -1)
	} else {
		fmt.Fprintln(ot, bfs(start))
	}
}
