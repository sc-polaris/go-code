package main

// Forward declaration of query API.
// func query(a int, b int) int{
//    ...
// }
// return int means matrix[x][y].

var matrix [][]int

func query(a int, b int) int {
	return matrix[a][b]
}

func getMinimumValue(N int) []int {
	const INF = int(1<<32 - 1)
	l, r := 0, N-1

	for l < r {
		mid := (l + r) >> 1

		k, val := 0, INF
		for i := 0; i < N; i++ {
			t := query(i, mid)
			if val > t {
				val = t
				k = i
			}
		}

		var left, right int
		if mid-1 >= 0 {
			left = query(k, mid-1)
		} else {
			left = INF
		}
		if mid+1 < N {
			right = query(k, mid+1)
		} else {
			right = INF
		}

		if val < left && val < right {
			return []int{k, mid}
		}
		if left < val {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	k, val := 0, INF
	for i := 0; i < N; i++ {
		t := query(i, r)
		if val > t {
			val = t
			k = i
		}
	}

	return []int{k, r}
}

func main() {
}
