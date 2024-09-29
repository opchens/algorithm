package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 转为数组
func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	i, j := 0, len(vals)-1
	for i <= j {
		if vals[i] != vals[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 递归
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	dummy := head
	var palindromeCheck func(node *ListNode) bool
	palindromeCheck = func(curhead *ListNode) bool {
		if curhead != nil {
			if !palindromeCheck(curhead.Next) {
				return false
			}
			if curhead.Val != dummy.Val {
				return false
			}
			dummy = dummy.Next
		}
		return true
	}
	return palindromeCheck(head)
}

// 双指针+反转
func isPalindrome3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	dummy := head
	midNode := GetMidPoint(head)
	reverNode := reverseList(midNode)
	for reverNode != nil || dummy != nil {
		if reverNode.Val != dummy.Val {
			return false
		}
		reverNode = reverNode.Next
		dummy = dummy.Next
	}
	return true

}

func GetMidPoint(node *ListNode) *ListNode {
	slow, fast := node, node
	for fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(node *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre, cur, nxt = nil, node, node.Next
	for cur != nil {
		cur.Next = pre
		pre = cur
		cur = nxt
		if nxt.Next != nil {
			nxt = nxt.Next
		}
	}
	return pre
}
