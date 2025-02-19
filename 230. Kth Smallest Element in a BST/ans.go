/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func visit(node *TreeNode, k int, list *[]int) bool {

    if node.Left != nil {
        success := visit(node.Left, k, list)
        if success {
            return true
        }
    }

    *list = append(*list, node.Val)
    if len(*list) == k {
        return true
    }

    if node.Right != nil {
        success := visit(node.Right, k, list)
        if success {
            return true
        }
    }

    return false
}

func kthSmallest(root *TreeNode, k int) int {
    list := []int{}
    visit(root, k, &list)

    return list[len(list) - 1]
}
