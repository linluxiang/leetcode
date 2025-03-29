import "container/list"

type Node [2]int // idx 0 is the city, idx 1 is the total vocation days

func maxVacationDays(flights [][]int, days [][]int) int {
    queue := list.New() 

    weeks := len(days[0])
    cities := len(days)

    week := 0
    queue.PushBack(Node{0, 0})
    for week < weeks {
        size := queue.Len()
        maxEachCity := map[int]int{}

        for range size {
            head := queue.Front().Value.(Node)
            queue.Remove(queue.Front())
            for target := range cities {
                if head[0] == target || flights[head[0]][target] == 1 {
                    if v, ok := maxEachCity[target]; !ok || v < head[1] + days[target][week] {
                        maxEachCity[target] = head[1] + days[target][week]
                    }
                }
            }
        }

        for city, day := range maxEachCity {
            queue.PushBack(Node{city, day})
        }

        week++
    }

    maxValue := 0
    for queue.Len() > 0 {
        head := queue.Front().Value.(Node)
        maxValue = max(maxValue, head[1])
        queue.Remove(queue.Front())
    }

    return maxValue
}
