package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{}
	p.Next = head
	x := FindFromEnd(p, n+1)
	x.Next = x.Next.Next
	return p.Next
}

func FindFromEnd(List *ListNode, k int) *ListNode {
	p1 := List
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	p2 := List
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
