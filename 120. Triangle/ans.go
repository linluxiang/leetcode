import "fmt"
var MAX = 99999
func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 {
        return 0
    }

    min := make([][]int, len(triangle))
    for i := range triangle {
        min[i] = make([]int, len(triangle[i]))
        for j := range min[i] {
            min[i][j] = MAX
        }
    }

    min[0][0] = triangle[0][0]
    for i := range len(triangle) - 1 {
        // fmt.Println(min[i])
        for j := range len(triangle[i]) {
            // fmt.Println(i, j)
            v := min[i][j] + triangle[i+1][j]
            if v < min[i+1][j] {
                min[i+1][j] = v
            }

            v = min[i][j] + triangle[i+1][j+1]
            if v < min[i+1][j+1] {
                min[i+1][j+1] = v
            }
        }
    }
    // fmt.Println(min[len(triangle) - 1])
    minV := 99999
    for _, v := range min[len(triangle) - 1] {
        if v < minV {
            minV = v
        }
    }
    return minV
}
