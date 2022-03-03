package main

import "fmt"

const N = 110
var res = [N][N]int{}
var n, m int
var dx = [4]int{0, 1, 0, -1}
var dy = [4]int{1, 0, -1, 0}

func main() {
    fmt.Scanln(&n, &m)

    for x, y, k, d := 0, 0, 1, 0; k <= m * n; k++ {
        res[x][y] = k
        a, b := x + dx[d], y + dy[d]
        if a < 0 || a >= n || b < 0 || b >= m || res[a][b] != 0 {
            d = (d + 1) % 4
            a, b = x + dx[d], y + dy[d]
        }
        x, y = a, b
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Printf("%d ", res[i][j])
        }
        fmt.Println()
    }
}