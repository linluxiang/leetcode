import "slices"
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
    Depth int
}

type ColNode struct {
    Val int
    Depth int
}

func Traverse(node *TreeNode, currCol int, depth int, column map[int][]*ColNode, minColumn *int)  {
    if node.Left != nil {
        Traverse(node.Left, currCol - 1, depth + 1, column, minColumn)
    }

    if _, ok := column[currCol]; !ok {
        // found new column
        column[currCol] = []*ColNode{}
        if currCol < *minColumn {
            *minColumn = currCol
        }
    }
    var i int = 0
    for i = 0; i < len(column[currCol]); i++ {
        itemDepth := column[currCol][i].Depth
        if depth < itemDepth {
            break
        }
        if depth >= itemDepth {
            continue
        }
    }

    column[currCol] = slices.Insert(column[currCol], i, &ColNode{node.Val, depth})

    if node.Right != nil {
        Traverse(node.Right, currCol + 1, depth + 1, column, minColumn)
    }
}

func verticalOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    result := [][]int{}
    column := map[int][]*ColNode{}
    minCol := 0
    Traverse(root, 0, 0, column, &minCol)

    for i := minCol; i < minCol + len(column); i++ {
        list := []int{}
        for _, v := range column[i] {
            list = append(list, v.Val)
        }
        result = append(result, list)
    }

    return result
}
