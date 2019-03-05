package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	nodeArr := make([]*ListNode, 0)
	for head != nil {
		nodeArr = append(nodeArr, head)
		head = head.Next
	}
	if n == 1 {
		if len(nodeArr) == 1 {
			return nil
		} else {
			nodeArr[len(nodeArr)-n-1].Next = nil
		}
	} else if n == len(nodeArr) {
		return nodeArr[len(nodeArr)-n].Next
	} else {
		nodeArr[len(nodeArr)-n-1].Next = nodeArr[len(nodeArr)-n+1]
	}
	return nodeArr[0]
}

func removeNthFromEnd_Best(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	fast := head
	slow := &ListNode{
		Next: head,
	}

	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	if slow.Next == head {
		return head.Next
	}

	slow.Next = slow.Next.Next

	return head
}
