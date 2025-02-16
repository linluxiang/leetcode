/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	p1 := l1
	p2 := l2

	ans := &ListNode{}

	pans := ans

	var value int
	var moreThanTen int = 0
	for p1 != nil && p2 != nil {
		value = p1.Val + p2.Val + moreThanTen
		moreThanTen = value / 10
		pans.Val = value % 10

		p1 = p1.Next
		p2 = p2.Next

		if p1 != nil || p2 != nil {
			pans.Next = &ListNode{}
			pans = pans.Next
		} else if p1 == nil && p2 == nil && moreThanTen != 0 {
            pans.Next = &ListNode{Val: moreThanTen}
			pans = pans.Next
        }
	}

	var unfinished *ListNode

	if p1 != nil {
		unfinished = p1
	}

	if p2 != nil {
		unfinished = p2
	}

	if unfinished == nil {
		return ans
	}

	for unfinished != nil {
		value = unfinished.Val + moreThanTen
		moreThanTen = value / 10
		pans.Val = value % 10
		unfinished = unfinished.Next
		if unfinished == nil && moreThanTen != 0 {
			pans.Next = &ListNode{ Val: moreThanTen }
			pans = pans.Next
            break
		} else if unfinished != nil {
            pans.Next = &ListNode{ }
			pans = pans.Next
        }

	}

	return ans
}
