package main

func moreThanHalfNumSolution(nums []int) int {
	value, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			value = v
			count++
		} else {
			if v == value {
				count++
			} else {
				count--
			}
		}
	}

	return value
}

func main() {

}
