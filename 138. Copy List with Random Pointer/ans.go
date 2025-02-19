/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
import "slices"

func copyRandomList(head *Node) *Node {
    list := []*Node{}
    newList := []*Node{}

    p := head
    var prev *Node = nil 
    for p != nil {
        list = append(list, p)
        newNode := &Node {
            Val: p.Val,
        }
        newList = append(newList, newNode)
        if prev == nil {
            prev = newNode
        } else {
            prev.Next = newNode
            prev = prev.Next
        }
        p = p.Next
    }

    if len(list) == 0 {
        return nil
    }

    indexes := make([]int, len(list))

    for i, item := range list {
        if item.Random == nil {
            indexes[i] = -1
        } else {
            indexes[i] = slices.Index(list, item.Random)
        }
    }

    for i, item := range newList {
        if indexes[i] != -1 {
            item.Random = newList[indexes[i]]
        }
    }

    return newList[0]
}
