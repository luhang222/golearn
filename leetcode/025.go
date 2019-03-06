package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	i := head
	j := head
	for count := 0; count < k-1; count++ {
		j = j.Next
		if j == nil {
			return i
		}
	}
	next := j.Next
	j.Next = nil
	reverse(i)
	i.Next = reverseKGroup(next, k)
	return j
}

func reverse(head *ListNode) *ListNode {
	i := head
	if i == nil || i.Next == nil {
		return i
	}
	h := reverse(i.Next)
	i.Next.Next = i
	i.Next = nil
	return h
}
