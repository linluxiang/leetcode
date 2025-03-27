/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

//  import "container/list"

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    if len(root.Children) == 0 {
        return 1
    }

    ans := 0

    for _, c := range root.Children {
        ans = max(ans, maxDepth(c) + 1)
    }

    return ans
}
