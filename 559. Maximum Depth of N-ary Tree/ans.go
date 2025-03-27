/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 import "container/list"

func maxDepth(root *Node) int {
    l := list.New()

    if root == nil {
        return 0
    }

    depth := 1

    l.PushBack(root)

    for l.Len() != 0 {
        size := l.Len()

        for i := 0; i < size; i++ {
            head := l.Front()
            l.Remove(head)
            node := head.Value.(*Node)
            for _, v := range node.Children {
                l.PushBack(v)
            }
        }

        depth ++
    }
    return depth - 1
}
