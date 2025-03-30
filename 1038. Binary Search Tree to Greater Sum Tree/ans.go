/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, toAdd *int)  {
    if root.Right != nil {
        traverse(root.Right, toAdd)
    }

    root.Val += *toAdd
    *toAdd = root.Val

    if root.Left != nil {
        traverse(root.Left, toAdd)
    }
}

func bstToGst(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Left == nil && root.Right == nil {
        return root
    }

    ans := 0
    traverse(root, &ans)

    return root
}
