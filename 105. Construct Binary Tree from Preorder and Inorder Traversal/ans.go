/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "slices"
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    if len(preorder) == 1 {
        return &TreeNode {
            Val: preorder[0],
        }
    }

    parentIdx := slices.Index(inorder, preorder[0])

    leftLength := parentIdx
    leftPreorder := preorder[1:leftLength + 1]
    leftInorder := inorder[:leftLength]

    rightPreorder := preorder[leftLength+1:]
    rightInorder := inorder[parentIdx+1:]
    
    parent := &TreeNode {
        Val: preorder[0],
    }
    left := buildTree(leftPreorder, leftInorder)
    right := buildTree(rightPreorder, rightInorder)

    parent.Left =left
    parent.Right = right
    return parent
}
