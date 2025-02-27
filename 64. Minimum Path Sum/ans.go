func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	if row == 1 && col == 1 {
		return grid[0][0]
	}

	for i := range row {
		for j := range grid[i] {
			if i == 0 && j >= 1 {
				grid[i][j] += grid[i][j-1]
			}

			if j == 0 && i >= 1 {
				grid[i][j] += grid[i-1][j]
			}
			if i >= 1 && j >= 1 {
				if grid[i-1][j] < grid[i][j-1] {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += grid[i][j-1]
				}
			}

		}
	}

	return grid[row-1][col-1]
}
