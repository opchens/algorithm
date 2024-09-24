package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 双指针
func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	p, q := dummy, head
	for q != nil {
		if q.Next != nil && q.Val == q.Next.Val {
			q = q.Next
			if q.Next == nil {
				p.Next = nil
			}
		} else {
			p.Next = q
			q = q.Next
			p = p.Next
		}
	}
	return dummy.Next
}
