package main

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	x := (tomatoSlices - 2*cheeseSlices) / 2
	y := (4*cheeseSlices - tomatoSlices) / 2
	if 4*x+2*y != tomatoSlices || x+y != cheeseSlices || x < 0 || y < 0 {
		return []int{}
	}
	return []int{x, y}
}

func main() {

}
