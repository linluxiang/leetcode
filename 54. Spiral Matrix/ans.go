type DIRECTION [2]int

var RIGHT DIRECTION = [2]int{0, 1}
var DOWN DIRECTION = [2]int{1, 0}
var LEFT DIRECTION = [2]int{0, -1}
var UP DIRECTION = [2]int{-1, 0}

var NEXT = map[DIRECTION]DIRECTION {
    RIGHT: DOWN,
    DOWN: LEFT,
    LEFT: UP,
    UP: RIGHT,
}

type POS [2]int

func spiralOrder(matrix [][]int) []int {
    row := len(matrix)
    col := len(matrix[0])

    visited := map[POS]bool{}

    direct := RIGHT
    curr := POS{0, 0}

    result := []int{}

    for len(visited) < row * col {
        // fmt.Println(curr, result, visited)
        result = append(result, matrix[curr[0]][curr[1]])
        visited[curr] = true
        x := curr[0] + direct[0]
        y := curr[1] + direct[1]

        if x < 0 || y < 0 || x >= row || y >= col || visited[POS{x, y}] == true {
            direct = NEXT[direct]
            x = curr[0] + direct[0]
            y = curr[1] + direct[1]
        }
        curr = POS{x, y}
    }

    return result
}
