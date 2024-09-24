package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := head
	hashkey := make(map[int]int)
	for cur != nil {
		if _, ok := hashkey[cur.Val]; !ok {
			hashkey[cur.Val] = 1
		} else {
			hashkey[cur.Val] += 1
		}
		cur = cur.Next
	}
	cur = head
	res := dummy
	for cur != nil {
		if hashkey[cur.Val] == 1 {
			dummy.Next = cur
			dummy = dummy.Next
		}
		cur = cur.Next
	}
	dummy.Next = nil
	return res.Next
}
