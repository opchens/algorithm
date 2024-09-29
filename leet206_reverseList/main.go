package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	t := &ListNode{Val: 2}
	t.Next = &ListNode{Val: 3}
	t.Next.Next = &ListNode{Val: 4}
	t.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println(reverseList(t))

}

func reverseList(head *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre, cur, nxt = nil, head, head.Next
	for head != nil {
		cur.Next = pre
		pre = cur
		cur = nxt
		if nxt != nil {
			nxt = nxt.Next
		}
	}
	return pre
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
