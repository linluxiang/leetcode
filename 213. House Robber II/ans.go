func maxRob(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    dp[1] = nums[1]
    for i := 2; i < len(nums); i++ {
        dp[i] = nums[i] + dp[i-2]
        if i - 3 >= 0 {
            dp[i] = max(dp[i], nums[i] + dp[i-3])
        }
    }

    return max(dp[len(nums)-1], dp[len(nums) - 2])
}

func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }

    return max(maxRob(nums[:len(nums) - 1]), maxRob(nums[1:len(nums)]))

}
