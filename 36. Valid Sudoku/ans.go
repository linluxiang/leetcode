
func isValid9Cells(cells []byte) bool {
	valid := make([]bool, 9)
	for _, c := range cells {
		if valid[c - 1] == true {
			return false
		}
        valid[c-1] = true
	}
	return true
}

func getRow(board [][]byte, row int) []byte {
	res := make([]byte, 0, 9)
	for _, c := range board[row] {
		if c == '.' {
			continue
		}
		res = append(res, c-byte('0'))
	}
	return res
}

func getCol(board [][]byte, col int) []byte {
	res := make([]byte, 0, 9)
	for i := 0; i < 9; i++ {
		c := board[i][col]
		if c == '.' {
			continue
		}
		res = append(res, c-byte('0'))
	}
	return res
}

func getBox(board [][]byte, x int, y int) []byte {
	res := make([]byte, 0, 9)
	points := make([]byte, 0, 9)
	points = append(points,
		board[x][y],
		board[x][y+1],
		board[x][y+2],
		board[x+1][y],
		board[x+1][y+1],
		board[x+1][y+2],
		board[x+2][y],
		board[x+2][y+1],
		board[x+2][y+2],
	)
	for _, c := range points {
		if c == '.' {
			continue
		}
		res = append(res, c-byte('0'))
	}
	return res
}

func isValidSudoku(board [][]byte) bool {

	for row := 0; row < 9; row++ {
		if isValid9Cells(getRow(board, row)) == false {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if isValid9Cells(getCol(board, col)) == false {
			return false
		}
	}

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if isValid9Cells(getBox(board, x*3, y*3)) == false {
				return false
			}
		}
	}

	return true

}
