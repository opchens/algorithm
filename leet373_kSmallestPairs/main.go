package main

import "container/heap"

func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	l := len(nums1)
	h := hp{nums1: nums1, nums2: nums2, data: nil}
	for i := 0; i < k && i < l; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	heap.Init(&h)
	for h.Len() > 0 && k > 0 {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		k--
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < len(nums2) {
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return ans
}

type pair struct {
	i, j int
}

type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.nums1) }
func (h hp) Less(i, j int) bool {
	a := h.data[i]
	b := h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}

func (h hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *hp) Pop() interface{} {
	old := h.data
	l := len(old)
	a := old[l-1]
	h.data = old[0 : l-1]
	return a
}

func (h *hp) Push(v interface{}) {
	h.data = append(h.data, v.(pair))
}
