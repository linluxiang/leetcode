/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */
import "fmt"
func tail(root *Node) *Node {
    if root == nil || root.Next == nil {
        return root
    }
    for root != nil {
        if root.Next == nil {
            return root
        }
        root = root.Next
    }

    return nil
}

func flatten(root *Node) *Node {
    head := root
    for head != nil {
        // fmt.Println("traverse:", head.Val)
        if head.Child != nil {
            child := flatten(head.Child)
            childTail := tail(child)
            next := head.Next
            head.Next = child
            child.Prev = head

            if next != nil {
                childTail.Next = next
                next.Prev = childTail
            }
            head.Child = nil
        }
        head = head.Next
    }

    return root
}
