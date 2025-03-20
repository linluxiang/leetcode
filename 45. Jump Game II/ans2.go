func jump(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    maxReach := 0
    lastMaxReach := 0    
    step := 0
    for i:=0; i < len(nums); i++ {
        // fmt.Println(i, nums[i], maxReach, lastMaxReach, step)
        if i + nums[i] > maxReach {
            maxReach = i+nums[i]
            if maxReach >= len(nums) - 1 {
                return step+1
            }
        }

        if i >= lastMaxReach {
            step += 1
            lastMaxReach = maxReach
        }

    }

    return step
}
