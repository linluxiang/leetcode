/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 type Entry struct {
    node *Node
    depth int
 }

func connect(root *Node) *Node {
	if root == nil {
        return nil
    }

    queue := []*Entry{&Entry{root, 0}}

    for len(queue) != 0 {
        head := queue[0]
        depth := head.depth
        // fmt.Println("head left: ", head.node.Left, " head right: ", head.node.Right)
        if head.node.Left != nil {
            queue = append(queue, &Entry{head.node.Left, depth+1})
        }
        if head.node.Right != nil {
            queue = append(queue, &Entry{head.node.Right, depth+1})
        }
        // fmt.Println("len: ", len(queue))
        if len(queue) == 1 {
            break
        }
        queue = queue[1:len(queue)]

        next := queue[0]
        // fmt.Println("head: ", head.node.Val, " next: ", next.node.Val, " queue: ", queue)

        for next.depth == head.depth {
            // fmt.Println("head: ", head.node.Val, " next: ", next.node.Val)
            head.node.Next = next.node
            head = next
            if next.node.Left != nil {
                queue = append(queue, &Entry{next.node.Left, depth+1})
            }
            if next.node.Right != nil {
                queue = append(queue, &Entry{next.node.Right, depth+1})
            }
            if len(queue) == 1 {
                break
            }
            queue = queue[1:len(queue)]

            next = queue[0]
        }

    }
    return root
}
