/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "slices"

func sortList(head *ListNode) *ListNode {
    list := make([]int, 0, 50000)
    p := head
    for p != nil {
        list = append(list, p.Val)
        p = p.Next
    }

    slices.Sort(list)

    h := &ListNode {

    }
    p = h
    for _, x := range list {
        newNode := &ListNode {
            Val: x,
        }
        p.Next = newNode
        p = p.Next
    }
    return h.Next
}
