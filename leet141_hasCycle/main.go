package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func hasCycle(head *ListNode) bool {
	dummy := &ListNode{}
	dummy.Next = head
	p := dummy.Next
	q := dummy.Next
	for p != nil && q != nil {
		if p == q {
			return true
		}
		p = p.Next
		q = q.Next.Next
	}
	return false
}
