package sol

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func isPossibleDivide(nums []int, k int) bool {
	nLen := len(nums)
	if k == 1 {
		return true
	}
	if nLen%k != 0 {
		return false
	}
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num] += 1
	}
	pq := MinHeap{}
	heap.Init(&pq)
	for key := range freq {
		heap.Push(&pq, key)
	}
	for pq.Len() > 0 {
		start := pq[0]
		end := start + k
		for num := start; num < end; num++ {
			val, ok := freq[num]
			if !ok {
				return false
			}
			val--
			freq[num] = val
			if val == 0 {
				heap.Pop(&pq)
			}
		}
	}
	return true
}
