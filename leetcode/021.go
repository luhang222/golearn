package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	m := &ListNode{}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := m
	for {
		if l1.Val < l2.Val {
			m.Next = l1
			l1 = l1.Next
			if l1 == nil {
				m.Next.Next = l2
				break
			}
		} else {
			m.Next = l2
			l2 = l2.Next
			if l2 == nil {
				m.Next.Next = l1
				break
			}
		}
		m = m.Next
	}
	return res.Next
}
