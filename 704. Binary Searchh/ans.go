func search(nums []int, target int) int {
    
    i,j := -1, len(nums)
    var mid int
    for i < j - 1 {
        mid = i + (j - i) / 2
        // fmt.Println(i, mid, j)
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            j = mid
        } else if nums[mid] < target {
            i = mid
        }

    }

    if i == -1 || j == len(nums) || (nums[i] != target && nums[j] != target) {
        return -1
    }
    return j
}
