package main

func maxScore(cardPoints []int, k int) int {
	sum, n := 0, len(cardPoints)
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	res := sum
	for i := 0; i < k; i++ {
		sum += cardPoints[n-i-1] - cardPoints[k-i-1]
		res = max(res, sum)
	}
	return res
}

func main() {

}
