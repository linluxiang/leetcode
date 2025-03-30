/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    sum := 0
    var traverse func(node *TreeNode) bool
    traverse = func(node *TreeNode) bool {
        defer func() {
            sum -= node.Val
        }()

        sum += node.Val

        if node.Left == nil && node.Right == nil {
            if sum == targetSum {
                return true
            }
        }

        if node.Left != nil {
            if traverse(node.Left) == true {
                return true
            }
        }

        if node.Right != nil {
            if traverse(node.Right) == true {
                return true
            }
        }

        return false

    }

    return traverse(root)
}
