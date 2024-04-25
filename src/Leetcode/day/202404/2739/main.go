package main

/*
	题干：

	卡车有两个油箱。给你两个整数，mainTank 表示主油箱中的燃料（以升为单位），additionalTank 表示副油箱中的燃料（以升为单位）。

	该卡车每耗费 1 升燃料都可以行驶 10 km。每当主油箱使用了 5 升燃料时，如果副油箱至少有 1 升燃料，则会将 1 升燃料从副油箱转移到主油箱。

	返回卡车可以行驶的最大距离。

	注意：从副油箱向主油箱注入燃料不是连续行为。这一事件会在每消耗 5 升燃料时突然且立即发生。

*/

func distanceTraveled(mainTank int, additionalTank int) (ans int) {
	for mainTank >= 5 {
		mainTank -= 5
		ans += 50
		if additionalTank > 0 {
			additionalTank--
			mainTank++
		}
	}
	return ans + mainTank*10
}

func distanceTraveled2(mainTank int, additionalTank int) (ans int) {
	for mainTank >= 5 {
		t := mainTank / 5
		ans += t * 50
		mainTank %= 5
		t = min(t, additionalTank)
		additionalTank -= t
		mainTank += t
	}
	return ans + mainTank*10
}
