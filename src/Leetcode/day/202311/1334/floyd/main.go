package floyd

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	const INF = math.MaxInt32 >> 1
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = INF
		}
	}

	for _, e := range edges {
		f, t, w := e[0], e[1], e[2]
		d[f][t], d[t][f] = w, w
	}

	for k := 0; k < n; k++ {
		d[k][k] = 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	ans, cnt := n, n+1
	for i := n - 1; i >= 0; i-- {
		t := 0
		for _, x := range d[i] {
			if x <= distanceThreshold {
				t++
			}
		}
		if t < cnt {
			cnt, ans = t, i
		}
	}

	return ans
}
