import "math"

func maxVacationDays(flights [][]int, days [][]int) int {
    weeks := len(days[0])
    cities := len(days)

    dpMap := make([][]int, cities)
    for i := range cities {
        dpMap[i] = make([]int, weeks)
        for j := range weeks {
            dpMap[i][j] = math.MaxInt
        }
    }

    var dfs func(city, week int) int
    dfs = func(city, week int) int { // the max value from city in week
        if week == weeks {
            return 0
        }
        if dpMap[city][week] != math.MaxInt {
            return dpMap[city][week]
        }
        maxV := 0
        for dest := range cities {
            if dest == city || flights[city][dest] == 1 {
                maxV = max(maxV, days[dest][week] + dfs(dest, week+1))
            }
        }
        dpMap[city][week] = maxV
        return maxV
    }

    maxValue := dfs(0, 0)

    return maxValue
}
