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
	return reverse(head)
}

func reverse(p *ListNode) *ListNode {
	fmt.Println(p)
	if p == nil || p.Next == nil {
		return p
	}
	j := reverse(p.Next)
	p.Next.Next = p
	p.Next = nil

	return j
}
