package main

func numberOfBoomerangs(points [][]int) (ans int) {
	cnt := make(map[int]int)
	for _, p1 := range points {
		clear(cnt)
		for _, p2 := range points {
			d2 := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
			ans += cnt[d2] * 2
			cnt[d2]++
		}
	}
	return ans
}

func main() {

}
