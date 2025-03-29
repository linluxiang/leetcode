import "fmt"

func rob(nums []int) int {
    lastIdx := len(nums) - 1
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = -1
    }

    var dfs func (house int) int

    dfs = func (house int) int {

        if house == lastIdx {
            return nums[lastIdx]
        }
        if dp[house] != -1 {
            return dp[house]
        }
        m := nums[house]

        if house + 3 <= lastIdx {
            m = max(m, nums[house] + dfs(house + 3))
        }

        if house + 2 <= lastIdx {
            m = max(m, nums[house] + dfs(house+2))

        }

        if house + 1 <= lastIdx {
            m = max(m, dfs(house+1))
        }
        dp[house] = m
        // fmt.Println(house, m)
        return m
    }

    return dfs(0)
}
