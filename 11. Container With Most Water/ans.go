import "slices"
import "fmt"

func min(a, b int) int {
    if a < b {
        return a 
    }
    return b
}

func volumn(height []int, i,j int) int {

    h := min(height[i], height[j])
    length := j - i
        // fmt.Println(i, j, h * length)
    return h * length
}

/*
这个解法的思路是这样的，对于一个固定的杆子来说，另外一个杆子在朝他靠近的过程中，只有找到更高的杆子才有可能有更大的面积。
*/
func maxArea(height []int) int {
    max := 0
    left := [][]int{}
    right := [][]int{}

    for i, v := range height {
        if len(left) == 0 {
            left = append(left, []int{i, v})
            continue
        }
        if v > left[len(left) - 1][1] {
            left = append(left, []int{i, v})
        }
    }

    // fmt.Println(left)

    for i := len(height) - 1; i > left[len(left) - 1][0]; i-- {
        v := height[i]
        if len(right) == 0 {
            right = append(right, []int{i, v})
            continue
        }
        if v > right[len(right) - 1][1] {
            right = append(right, []int{i, v})
        }
    }

    slices.Reverse(right)

    // fmt.Println(right)


    for _, l := range left {
        v := volumn(height, l[0], left[len(left) - 1][0])

        if v > max {
            max = v
        }
    }

    for _, r := range right{
        v := volumn(height, r[0], right[len(right) - 1][0])
        if v > max {
            max = v
        }
    }

    for _, l := range left {
        for _, r := range right {
            v := volumn(height, l[0], r[0])
            if v > max {
                max = v
            }
        }
    }

    return max
}

