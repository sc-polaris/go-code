package main

type FrequencyTracker struct {
	cnt  map[int]int // number 的出现次数
	freq map[int]int // number 的出现次数的出现次数
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{make(map[int]int), make(map[int]int)}
}

func (f *FrequencyTracker) Add(number int) {
	f.freq[f.cnt[number]]-- // 去掉一个旧的 cnt[number]
	f.cnt[number]++
	f.freq[f.cnt[number]]++ // 添加一个新的 cnt[number]
}

func (f *FrequencyTracker) DeleteOne(number int) {
	if f.cnt[number] > 0 {
		f.freq[f.cnt[number]]-- // 去掉一个旧的 cnt[number]
		f.cnt[number]--
		f.freq[f.cnt[number]]++ // 添加一个新的 cnt[number]
	}
}

func (f *FrequencyTracker) HasFrequency(frequency int) bool {
	return f.freq[frequency] > 0 // 至少有一个 number 的出现次数恰好为 frequency
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */

func main() {

}
