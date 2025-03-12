/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{} // an empty head
	newTail := newHead

	p := head
	for p != nil {
		if p.Next == nil {
			newTail.Next = p
			newTail = newTail.Next
			break
		}

		next := p.Next
		curr := p
		p = next.Next
		next.Next = curr
        curr.Next = nil
		newTail.Next = next
		newTail = curr
	}
	return newHead.Next
}
