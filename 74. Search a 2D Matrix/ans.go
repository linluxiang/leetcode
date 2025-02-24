import "slices"
import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    row := len(matrix)
    col := len(matrix[0])

    data := make([]int, row * col)
    for i := range row {
        for j := range col {
            // fmt.Println(i*row+j)
            data[i * col + j] = matrix[i][j]
            // data = append(data, matrix[i][j])
        }
    }

    // fmt.Printf("%v\n", data)

    _, found := slices.BinarySearch(data, target)
    return found
}
