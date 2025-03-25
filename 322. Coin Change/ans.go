// import "slices"
// import "math"
// f(x) == min(f(x-c0), f(x-c1)....f(x-cn)) + 1 // if valid 
// or -1 // if not valid
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    slices.SortFunc(coins, func(a ,b int) int {
        return b - a
    })
    dpMap := map[int]int{}

    var dfs func (target int) int
    dfs = func (target int) int {
        if dpMap[target] != 0 {
            return dpMap[target]
        }

        if target == 0 {
            return 0
        }

        if target < 0 {
            return -1
        }

        minValue := math.MaxInt
        for _, coin := range coins {
            if coin > target {
                continue
            }

            v := dfs(target - coin)
            if v >= 0 && v < minValue {
                minValue = v
            }
        }

        if minValue == math.MaxInt {
            // no valid solution
            dpMap[target] = -1
        } else {
            dpMap[target] = minValue + 1

        }

        return dpMap[target]
    }

    dfs(amount)

    if dpMap[amount] == 0 {
        return -1
    }
    return dpMap[amount]
}

