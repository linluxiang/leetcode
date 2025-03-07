/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var evenHead, oddHead, evenTail, oddTail *ListNode

    // create head node for convenience
    evenHead, oddHead = &ListNode{}, &ListNode{}
    evenTail, oddTail = evenHead, oddHead

    curr := head
    cnt := 1
    for curr != nil {
        next := curr.Next
        curr.Next = nil
        if cnt % 2  == 0 {
            evenTail.Next = curr
            evenTail = evenTail.Next
        } else {
            oddTail.Next = curr
            oddTail = oddTail.Next
        }
        cnt++
        curr = next
    }

    oddTail.Next = evenHead.Next
    return oddHead.Next
}
