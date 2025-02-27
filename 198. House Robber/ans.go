import "fmt"

func rob(nums []int) int {
    dp := make([]int, len(nums))

    if len(nums) == 1 {
        return nums[0]
    }
    dp[0] = nums[0]
    dp[1] = nums[1]
    max := 0
    for i := range nums {
        if i - 3 >= 0 {
            if dp[i-2] > dp[i-3] {
                dp[i] = dp[i-2] + nums[i]
            } else {
                dp[i] = dp[i-3] + nums[i]
            }
        } else if i - 2 >= 0{
            dp[i] = dp[i-2] + nums[i]
        } else {
            dp[i] = nums[i]
        }
        if dp[i] > max {
            max = dp[i]
        }
    }
    //fmt.Println(dp)
    return max
}
