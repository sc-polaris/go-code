package main

/*
	给你一份旅游线路图，该线路图中的旅行线路用数组 paths 表示，其中 paths[i] = [cityAi, cityBi] 表示该线路将会从 cityAi
	直接前往 cityBi 。请你找出这次旅行的终点站，即没有任何可以通往其他城市的线路的城市。

	题目数据保证线路图会形成一条不存在循环的线路，因此恰有一个旅行终点站。
*/

func destCity(paths [][]string) string {
	m := make(map[string]bool)
	for _, p := range paths {
		m[p[0]] = true
	}
	for _, p := range paths {
		if !m[p[1]] {
			return p[1]
		}
	}
	return ""
}
