package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	array1 := []int{}
	array2 := []int{}
	for l1 != nil {
		array1 = append(array1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		array2 = append(array2, l2.Val)
		l2 = l2.Next
	}
	dummy := &ListNode{-1, nil}

	carry := 0
	for len(array1) > 0 || len(array2) > 0 || carry > 0 {
		res := carry
		if len(array1) > 0 {
			res += array1[len(array1)-1]
			array1 = array1[0 : len(array1)-1]
		}
		if len(array2) > 0 {
			res += array2[len(array2)-1]
			array2 = array2[0 : len(array2)-1]
		}

		carry = res / 10
		res = res % 10
		cur := &ListNode{res, dummy.Next}
		dummy.Next = cur
	}
	return dummy.Next
}
