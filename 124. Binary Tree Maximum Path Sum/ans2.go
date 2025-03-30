import "math"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 返回值是最大的单路path，要么左+自己，要么右+自己，要么自己，但是全局判断的时候是左加右
func maxPathSum2(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }

    leftMaxSum := maxPathSum2(root.Left, ans)

    rightMaxSum := maxPathSum2(root.Right, ans)
    newans := max(*ans, root.Val, root.Val + leftMaxSum, root.Val + rightMaxSum, root.Val + leftMaxSum + rightMaxSum)
    *ans = newans

    return max(root.Val, root.Val + leftMaxSum, root.Val + rightMaxSum)
}

func maxPathSum(root *TreeNode) int {
    ans := math.MinInt
    maxPathSum2(root, &ans)
    return ans
}
