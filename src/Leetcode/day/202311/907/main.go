package main

func sumSubarrayMins(arr []int) (res int) {
	arr = append(arr, -1)
	st := []int{-1} // 哨兵
	for r, x := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= x {
			i := st[len(st)-1]
			st = st[:len(st)-1]
			res += arr[i] * (i - st[len(st)-1]) * (r - i)
		}
		st = append(st, r)
	}
	return res % (1e9 + 7)
}
func main() {

}
