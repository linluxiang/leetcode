/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reserve(head *ListNode) *ListNode {
    var prev *ListNode 
    next := head
    for next != nil {
        nextNext := next.Next
        next.Next = prev
        prev = next
        next = nextNext
    }
    return prev
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var prev, next *ListNode
    cnt := 1
    if head == nil {
        return head
    }
    p := head
    for p != nil {
        if cnt == left - 1 {
            prev = p
        }

        if cnt == right {
            next = p.Next
            p.Next = nil
        }

        p = p.Next
        cnt++
    }

    var start *ListNode
    if prev == nil {
        start = head
    } else {
        start = prev.Next
        prev.Next = nil
    }


    newHead := reserve(start)

    if prev != nil {
        prev.Next = newHead
    }

    p = newHead
    for p != nil {
        if p.Next == nil {
            break
        }
        p = p.Next
    }
    p.Next = next
    if prev != nil {
        return head
    }
    return newHead
}
