/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Node struct {
    TreeNode *TreeNode
    Depth int
 }

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    result := [][]int{}
    levels := map[int][]int{}
    queue := []*Node{&Node{root, 0}}

    for len(queue) > 0 {
        head := queue[0]
        if _, ok := levels[head.Depth]; !ok {
            levels[head.Depth] = []int{}
        }

        levels[head.Depth] = append(levels[head.Depth], head.TreeNode.Val)

        if head.TreeNode.Left != nil {
            queue = append(queue, &Node{head.TreeNode.Left, head.Depth + 1})
        }

        if head.TreeNode.Right != nil {
            queue = append(queue, &Node{head.TreeNode.Right, head.Depth + 1})
        }

        queue = queue[1:len(queue)]

    }

    for i := range len(levels) {
        result = append(result, levels[i])
    }

    return result
}
