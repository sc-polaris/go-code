package main

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

}
