package main

import "container/heap"

func kthSmallest(matrix [][]int, k int) int {
	pq := &hp{}
	heap.Init(pq)

	for i := range matrix {
		heap.Push(pq, []int{matrix[i][0], i, 0})
	}
	res := -1
	for k > 0 {
		cur := heap.Pop(pq).([]int)
		res = cur[0]
		k--
		i, j := cur[1], cur[2]
		if j+1 < len(matrix[i]) {
			heap.Push(pq, []int{matrix[i][j+1], i, j + 1})
		}
	}
	return res
}

type hp [][]int

func (h hp) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h hp) Len() int           { return len(h) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *hp) Pop() interface{} {
	tmp := *h
	l := len(tmp)
	x := tmp[l-1]
	*h = tmp[0 : l-1]
	return x
}
