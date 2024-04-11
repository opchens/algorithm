package main

// 难点在于如何快速找到k个中最小的节点，接到结果链表上

//分治

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(list []*ListNode, l, r int) *ListNode {
	if l == r {
		return list[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return merageTwoListNode(merge(list, l, mid), merge(list, mid+1, r))
}

func merageTwoListNode(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	dummyHead := &ListNode{}
	p := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummyHead.Next
}
