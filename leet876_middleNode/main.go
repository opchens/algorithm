package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			t := head
			for t != slow {
				t = t.Next
				slow = slow.Next
			}
			return t
		}
	}
	return nil
}
