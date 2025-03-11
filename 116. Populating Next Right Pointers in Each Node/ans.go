/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func tryconnect(root *Node){
    if root == nil {
        return
    }
    if root.Left != nil {
        root.Left.Next = root.Right
    }
    if root.Next != nil && root.Right != nil {
        root.Right.Next = root.Next.Left
    }
    tryconnect(root.Left)
    tryconnect(root.Right)
}

func connect(root *Node) *Node {
    tryconnect(root)
    return root
}
