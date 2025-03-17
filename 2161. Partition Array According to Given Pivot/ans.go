func pivotArray(nums []int, pivot int) []int {
    if len(nums) == 0 || len(nums) == 1{
        return nums
    }

    
    smallerIdx, biggerIdx := 0, len(nums) - 1

    ans := make([]int, len(nums))

    for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
        if nums[i] < pivot {
            ans[smallerIdx] = nums[i]
            smallerIdx += 1
        }
        if nums[j] > pivot {
            ans[biggerIdx] = nums[j]
            biggerIdx -= 1
        }
    }

    for i := smallerIdx; i <= biggerIdx; i ++ {
        ans[i] = pivot
    }
    return ans
}
