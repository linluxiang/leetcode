// import "slices"
import "fmt"

// func min(a, b int) int {
//     if a < b {
//         return a 
//     }
//     return b
// }

func volumn(height []int, i,j int) int {

    h := min(height[i], height[j])
    length := j - i
        // fmt.Println(i, j, h * length)
    return h * length
}

/*
这个O(n)的解法是从左和右两边往中间收缩，收缩的时候只移动矮的那个边。
因为面积是由短的那个边决定的，如果是移动高的，那么很可能会遇到矮的导致容量减少，遇到高的容积也不会增加。
如果移动矮的，遇到矮的容量减少，遇到高的容量增加。
所以两边靠近的过程中，应该从矮的那个往高的那个移动。
*/
func maxArea(height []int) int {
    max := 0
    
    i, j := 0, len(height) - 1
    for i != j {
        v := volumn(height, i, j)
        if v > max {
            max = v
        }

        if min(height[i], height[j]) == height[i] {
            i++
        } else {
            j--
        }
    }
    return max
}
