func findMin(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    i, j := 0, len(nums) - 1

    for i < j - 1 {
        mid := i + (j - i) >> 1
        
        // fmt.Println(i, mid, j)
        if nums[i] <= nums[mid] && nums[mid] <= nums[j] {
            return nums[i]
        }

        if nums[mid] > nums[j] {
            i = mid
            continue
        }

        if nums[mid] < nums[j] {
            j = mid
            continue
        }
    }

    return min(nums[i], nums[j])
}
