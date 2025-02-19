/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func visit(node *TreeNode, depth int, result *[]int, marked map[int]bool) {
    if _, ok := marked[depth]; !ok {
        *result = append(*result, node.Val)
        marked[depth] = true
    }

    if node.Right != nil {
        visit(node.Right, depth + 1, result, marked)
    }

    if node.Left != nil {
        visit(node.Left, depth + 1, result, marked)
    }
}


func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }

    result := []int{}
    marked := map[int]bool{}
    visit(root, 1, &result, marked)

    return result
}
