import "slices"

func twoCitySchedCost(costs [][]int) int {
    // 根据一个人去a-去b的差排序，选择前n个最小的人去a，其他人去b
    // 贪心
    slices.SortFunc(costs, func (a, b []int) int {
        return a[0] - a[1] - (b[0] - b[1])
    })

    n := len(costs) / 2
    a := costs[:n]
    b := costs[n:]

    sum := 0 
    for _, v := range a {
        sum += v[0]
    }

    for _, v := range b {
        sum += v[1]
    }

    return sum
}
