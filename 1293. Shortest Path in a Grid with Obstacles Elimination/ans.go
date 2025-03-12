import "fmt"

var MAX = 1000

var DIR = [4][2]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

type Node struct {
	X     int
	Y     int
	Depth []int // Depth[i] means the shortest distance it can use by eliminate i blocks
}

func shortestPath(grid [][]int, k int) int {
	queue := []*Node{}
	nodeInfo := map[[2]int]*Node{}
	rows := len(grid)
	cols := len(grid[0])
	queue = append(queue, &Node{0, 0, make([]int, k+1)})
	for i := range k + 1 {
		queue[0].Depth[i] = 0
	}

	nodeInfo[[2]int{0, 0}] = queue[0]

	for len(queue) != 0 {
		h := queue[0]
		queue = queue[1:]

		curr := grid[h.X][h.Y]
		currNode := nodeInfo[[2]int{h.X, h.Y}]
		// normal grid
		for _, dir := range DIR {
			x := h.X + dir[0]
			y := h.Y + dir[1]
            if x == 0 && y == 0 {
                continue
            }
			if x < 0 || y < 0 || x >= rows || y >= cols {
				continue
			}
            pos := [2]int{x, y}
			if v := nodeInfo[pos]; v == nil {
				nodeInfo[pos] = &Node{x, y, make([]int, k+1)}
                for i := range k + 1 {
		            nodeInfo[pos].Depth[i] = MAX
	            }  
			}
			node := nodeInfo[pos]

            nextGrid := grid[x][y]

			valueUpdated := false

			for i := range k + 1 {
                if curr == 0 && nextGrid == 0 {
                    // from normal to normal

                    if (node.Depth[i] > currNode.Depth[i]+1) {

					    node.Depth[i] = currNode.Depth[i] + 1
					    valueUpdated = true
				    }
                } else if curr == 0 && nextGrid == 1 {
                    node.Depth[0] = MAX
                    if i + 1 >= len(node.Depth) {
                        break
                    }
                    if (node.Depth[i+1] > currNode.Depth[i] + 1) {
                        node.Depth[i+1] = currNode.Depth[i] + 1
                        valueUpdated = true
                    }
                } else if curr == 1 && nextGrid == 0 {
                    if i == 0 {
                        continue
                    }
                    if (node.Depth[i] > currNode.Depth[i] + 1) {
                        node.Depth[i] = currNode.Depth[i] + 1
                        valueUpdated = true
                    }
                } else if curr == 1 && nextGrid == 1 {
                    node.Depth[0] = MAX
                    if i + 1 >= len(node.Depth) {
                        break
                    }
                    if (node.Depth[i+1] > currNode.Depth[i] + 1) {
                        node.Depth[i+1] = currNode.Depth[i] + 1
                        valueUpdated = true
                    }
                }
			}
			if valueUpdated {
				queue = append(queue, node)
			}
		}
	}

    // for m := range rows {
    //     r := []int{}
    //     for n := range cols {
    //         node := nodeInfo[[2]int{m, n}]
    //         //fmt.Println(m, n, node.Depth)
    //         if node != nil {
    //             r = append(r, node.Depth[0])

    //         }
    //     }
    //     fmt.Println(r)
    // }

	lastPos := [2]int{rows - 1, cols - 1}
        // fmt.Println(nodeInfo[lastPos].Depth)

	if nodeInfo[lastPos] == nil {
		return -1
	}

	minV := rows * cols
	for _, v := range nodeInfo[lastPos].Depth {
		if v < minV {
			minV = v
		}
	}
	if minV == rows*cols {
		return -1
	}
	return minV
}
