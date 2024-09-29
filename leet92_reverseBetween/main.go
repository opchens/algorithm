package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	for i := 0; i < left; i++ {
		pre = pre.Next
	}
	leftPoint := pre.Next
	cur := leftPoint
	for i := 0; i < left-right+1; i++ {
		cur = cur.Next
	}
	rightPoint := cur.Next
	cur.Next = nil
	pre.Next = nil
	reverseLinkList(leftPoint)
	pre.Next = cur
	leftPoint.Next = rightPoint
	return dummy.Next

}

func reverseLinkList(head *ListNode) {
	prev, curr := &ListNode{}, head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}
