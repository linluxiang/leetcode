/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var prev *ListNode 
    next := head
    for next != nil {
        // fmt.Println("next: ", next, " prev: ", prev)
        nextNext := next.Next
        next.Next = prev
        prev = next
        next = nextNext
    }

    return prev
}
