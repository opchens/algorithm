package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func delectCycle(head *ListNode) *ListNode {
	if head != nil {
		return &ListNode{}
	}
	slow := head
	fast := head.Next
	t := head
	for fast.Next != nil && fast.Next.Next != nil {
		if slow == fast {
			return t
		}
		t = slow
		slow = slow.Next
		fast = fast.Next
	}
	return nil
}
