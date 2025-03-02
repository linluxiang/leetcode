// import "slices"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Node struct {
    Val *TreeNode
    Col int
}


func verticalOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    result := [][]int{}
    columns := map[int][]*Node{}
    minCol := 0

    queue := []*Node{ &Node{root, 0}}

    for len(queue) > 0 {
        head := queue[0]
        if head.Col < minCol {
            minCol = head.Col
        }

        if _, ok := columns[head.Col]; !ok {
            columns[head.Col] = []*Node{}
        } 

        columns[head.Col] = append(columns[head.Col], head)

        if head.Val.Left != nil {
            queue = append(queue, &Node{head.Val.Left, head.Col - 1})
        }

        if head.Val.Right != nil {
            queue = append(queue, &Node{head.Val.Right, head.Col + 1})
        }

        queue = queue[1:len(queue)]
    }

    for i := minCol; i < minCol + len(columns); i ++ {
        l := []int{}
        for _, v := range columns[i] {
            l = append(l, v.Val.Val)
        }
        result = append(result, l)
    }

    return result
}
