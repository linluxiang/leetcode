import "math"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 最大的单路path，要么左+自己，要么右+自己，要么自己
func maxPathSum2(root *TreeNode, ans *int) int {
    result := root.Val

    if root.Val > *ans {
        *ans = root.Val
    }

    if root.Left != nil && root.Right == nil {
        leftMaxSum := maxPathSum2(root.Left, ans)
        newans := max(*ans, root.Val + leftMaxSum)
        *ans = newans
        return max(root.Val, root.Val + leftMaxSum)
    }

    if root.Right != nil && root.Left == nil {
        rightMaxSum := maxPathSum2(root.Right, ans)
        newans := max(*ans, root.Val + rightMaxSum)
        *ans = newans
        return max(root.Val, root.Val + rightMaxSum)
    }

    if root.Left !=nil && root.Right != nil {
        leftMaxSum := maxPathSum2(root.Left, ans)
        rightMaxSum := maxPathSum2(root.Right, ans)
        newans := max(*ans, root.Val, root.Val + rightMaxSum, root.Val+leftMaxSum, root.Val+leftMaxSum+rightMaxSum)
        *ans = newans
        return max(root.Val, root.Val + leftMaxSum, root.Val + rightMaxSum)
    }

    // return max(root.Val, root.Val+leftMaxSum0, root.Val+rightMaxSum0, root.Val+leftMaxSum0+rightMaxSum0), max(leftMaxSum1, rightMaxSum1, leftMaxSum0, rightMaxSum0)
    return result
}

func maxPathSum(root *TreeNode) int {
    ans := math.MinInt
    maxPathSum2(root, &ans)
    return ans
}
