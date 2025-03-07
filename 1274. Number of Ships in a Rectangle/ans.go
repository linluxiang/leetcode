/**
 * // This is Sea's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Sea struct {
 *     func hasShips(topRight, bottomLeft []int) bool {}
 * }
 */
import "fmt"

 func count(sea Sea, topRight, bottomLeft []int) int {
    if !sea.hasShips(topRight, bottomLeft) {
        return 0
    }
    x1, x2 := bottomLeft[0], topRight[0]
    y1, y2 := bottomLeft[1], topRight[1]

    if x1 == x2 && y1 == y2 {
        return 1
    }

    xmid := x1 + (x2-x1) >> 1
    ymid := y1 + (y2-y1) >> 1
    cnt := 0

    cnt += count(sea, []int{xmid, ymid}, []int{x1, y1})
    if xmid + 1 <= x2 {
        cnt += count(sea, []int{x2, ymid}, []int{xmid +1, y1})
    }
    if ymid + 1 <= y2 {
        cnt += count(sea, []int{xmid, y2}, []int{x1, ymid+1})
    }
    if xmid + 1 <= x2 && ymid+1 <= y2 {
        cnt += count(sea, []int{x2, y2}, []int{xmid+1, ymid+1})
    }
    // fmt.Println("cnt: ", cnt)

    return cnt
 }

func countShips(sea Sea, topRight, bottomLeft []int) int {
    return count(sea, topRight, bottomLeft)
}
