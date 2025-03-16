var Directions = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Pos [2]int

type Entry struct {
	pos Pos
	len int
}

func dfs(pos Pos, length int, rows, cols int, visited map[Pos]bool, board [][]byte, word string) bool {
	defer func() {
		visited[pos] = false
	}()

	if length == len(word) {
		return true
	}

	visited[pos] = true

	for _, d := range Directions {
		x := pos[0] + d[0]
		y := pos[1] + d[1]
		newPos := Pos{x, y}

		if x < 0 || y < 0 || x >= rows || y >= cols || visited[newPos] == true || board[x][y] != word[length] {
			continue
		}

		v := dfs(newPos, length+1, rows, cols, visited, board, word)
		if v == true {
			return true
		}
	}
    return false
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
    if len(word) > rows * cols {
        return false
    }

    appears := map[byte]int{}

	startPos := []Pos{}
	for i := range rows {
		for j := range cols {
			if board[i][j] == word[0] {
				startPos = append(startPos, Pos{i, j})
			}
            appears[board[i][j]] += 1
		}
	}

    appearsInWord := map[byte]int{}

    for _, c := range []byte(word) {
        appearsInWord[c] += 1
    }

    for c, v := range appears {
        if v < appearsInWord[c] {
            return false
        }
    }

	for _, pos := range startPos {
		visited := map[Pos]bool{}
		v := dfs(pos, 1, rows, cols, visited, board, word)
		if v == true {
			return true
		}
	}

	return false
}
