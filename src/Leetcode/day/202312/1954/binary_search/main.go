package main

import (
	"sort"
)

func minimumPerimeter(neededApples int64) int64 {
	ans := sort.Search(100000, func(i int) bool {
		return int64(2*i*(i+1)*(2*i+1)) >= neededApples
	})

	return 8 * int64(ans)
}

func main() {

}
