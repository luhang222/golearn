package leetcode

func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	i := head
	j := i.Next
	if j != nil {
		k := i.Next.Next
		i.Next.Next = i
		if k != nil {
			i.Next = SwapPairs(k)
		} else {
			i.Next = k
		}
		return j
	} else {
		return i
	}
}
