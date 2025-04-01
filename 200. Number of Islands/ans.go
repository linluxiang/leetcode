//import "container/list"

type Point struct {
	X, Y int
}


func expand(grid [][]byte, i,j, islandId int, islands map[Point]int) {

        curr := Point {i, j}
        islands[curr] = islandId

        if curr.X - 1 >= 0 && grid[curr.X - 1][curr.Y] == '1' && islands[Point{curr.X - 1, curr.Y}] == 0 {
            expand(grid, i - 1, j, islandId, islands)
        }

        if curr.X + 1 < len(grid) && grid[curr.X + 1][curr.Y] == '1' && islands[Point{curr.X + 1, curr.Y}] == 0 {
            expand(grid, i + 1, j, islandId, islands)
        }

        if curr.Y - 1 >=0 && grid[curr.X][curr.Y - 1] == '1' && islands[Point{curr.X, curr.Y - 1}] == 0 {
            expand(grid, i, j - 1, islandId, islands)
        }

        if curr.Y + 1 < len(grid[curr.X]) && grid[curr.X][curr.Y + 1] == '1' && islands[Point{curr.X, curr.Y+1}] == 0 {
            expand(grid, i, j + 1, islandId, islands)
        }
}

func numIslands(grid [][]byte) int {

	var islands = make(map[Point]int)

	islandId := 0
	for i, row := range grid {
		for j, value := range row {
			p := Point{i, j}
			if value == '0' {
				continue
			}
            
			if _, ok := islands[p]; !ok {
				// new island
				islandId += 1
				expand(grid, i, j, islandId, islands)
				continue
			}
		}
	}
	return islandId
}
