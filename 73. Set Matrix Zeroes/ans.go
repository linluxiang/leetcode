type Pos [2]int


func setZeroes(matrix [][]int)  {
    rows := len(matrix)
    cols := len(matrix[0])

    zeroRow := map[int]bool{}
    zeroCol := map[int]bool{}

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == 0 {
            zeroRow[i] = true
            zeroCol[j] = true
            }

        }
    }

    for row := range zeroRow {
        for i := 0; i < cols; i++ {
            matrix[row][i] = 0
        }
    }

    for col := range zeroCol {
        for i := 0; i < rows; i++ {
            matrix[i][col] = 0
        }
    }

}
