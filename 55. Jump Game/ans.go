func canJump(nums []int) bool {
    maxCanReach := 0
    for i, v := range nums {
        if i == len(nums) - 1 {
            break
        }
        if i > maxCanReach {
            break
        }
        if i + v > maxCanReach {
            maxCanReach = i + v
        }
        // fmt.Println(i, maxCanReach)
    }
    if maxCanReach >= len(nums) - 1 {
        return true
    }
    return false
}
