type Point struct {
    X,Y int
}

func expandFromO(board [][]byte, i, j int, visited map[Point]bool) {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
        return
    }

    visited[Point{i,j}] = true
    // left
    if i - 1 >= 0 && board[i - 1][j] == 'O' && visited[Point{i-1, j}] == false {
        expandFromO(board, i-1, j, visited)
    }

    // right
    if i + 1 < len(board) && board[i + 1][j] == 'O' && visited[Point{i+1, j}] == false {
        expandFromO(board, i+1, j, visited)
    }

    // up
    if j - 1 >= 0 && board[i][j-1] == 'O' && visited[Point{i, j-1}] == false {
        expandFromO(board, i, j-1, visited)
    }

    // down
    if j + 1 < len(board[0]) && board[i][j + 1] == 'O' && visited[Point{i, j+1}] == false {
        expandFromO(board, i, j+1, visited)
    }
}

func solve(board [][]byte)  {
    if len(board) == 0 {
        return
    }

    visited := map[Point]bool{}

    for i := range board {
        for j := range board[i] {
            if (i == 0 || i == len(board) - 1 || j == 0 || j == len(board[0]) - 1) && board[i][j] == 'O' {
                expandFromO(board, i, j, visited)
            }
        }
    }

    for i := range board {
        for j := range board[i] {
            if !visited[Point{i, j}] {
                board[i][j] = 'X'
            }
        }
    }
}
