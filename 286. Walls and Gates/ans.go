import "fmt"

var INF int = 2147483647

var dirs = [4][2]int {
    {-1, 0},
    {0, -1},
    {1, 0},
    {0, 1},
}

func wallsAndGates(rooms [][]int)  {
    rows := len(rooms)
    cols := len(rooms[0])

    queue := make([][]int, 0, rows * cols)
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if rooms[i][j] == 0 {
                // gate
                queue = append(queue, []int{i, j})

            }
        }
    }

    for len(queue) > 0 {
        head := queue[0]
        queue = queue[1:]

        i, j := head[0], head[1]
        //rooms[i][j] = dis
        // fmt.Println(i, j, dis)

        for _, d := range dirs {
            ni, nj := i + d[0], j + d[1]
            // 这里是关键的优化点，因为是从所有的gate开始的bfs，所以所有距离门为n的格子是同时访问的，如果某个格子已经有值了，那么一定是最近的。
            if ni < 0 || nj < 0 || ni >= rows || nj >= cols || rooms[ni][nj] != INF {
                continue
            }
            rooms[ni][nj] = rooms[i][j] + 1
            queue = append(queue, []int{ni, nj})
        }

        // // check up
        // if i - 1 >= 0 && rooms[i-1][j] != 0 && rooms[i-1][j] != -1  && rooms[i-1][j] == INF {
        //     rooms[i-1][j] = rooms[i][j] + 1
        //     queue = append(queue, []int{i-1, j})
        // }
        // // check left
        // if j - 1 >= 0 && rooms[i][j-1] != 0 && rooms[i][j-1] != -1 && rooms[i][j-1] == INF{
        //     rooms[i][j-1] = rooms[i][j] + 1
        //     queue = append(queue, []int{i, j-1})
        // }
        // // check down
        // if i + 1 < rows && rooms[i+1][j] != 0 && rooms[i+1][j] != -1 && rooms[i+1][j] == INF {
        //     rooms[i+1][j] = rooms[i][j] + 1
        //     queue = append(queue, []int{i+1, j})
        // }
        // // check right
        // if j + 1 < cols && rooms[i][j+1] != 0 && rooms[i][j+1] != -1 && rooms[i][j+1] == INF {
        //     rooms[i][j+1] = rooms[i][j] + 1
        //     queue = append(queue, []int{i, j+1})
        // }
    }
}

