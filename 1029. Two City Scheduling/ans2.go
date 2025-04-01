import "slices"


// 这个算法的核心思想是，优先处理去a城市和b城市之间的差值比较大的
func twoCitySchedCost(costs [][]int) int {
    slices.SortFunc(costs, func (a, b []int) int {
        return int(math.Abs(float64(b[0] - b[1]))) - int(math.Abs(float64(a[0] - a[1])))
    })

    cost := 0
    n := len(costs) / 2
    city1People := 0
    city2People := 0
    for i := 0; i < len(costs); i++ {
        c := costs[i]
        if city1People == n {
            city2People += 1
            cost += c[1]
            continue
        }

        if city2People == n {
            city1People += 1
            cost += c[0]
            continue
        }

        if c[0] < c[1] {
            city1People += 1
            cost += c[0]
        } else {
            city2People += 1
            cost += c[1]
        }
    }

    return cost
}
