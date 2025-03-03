func subarraySum(nums []int, k int) int {
    if len(nums) == 1 {
        if nums[0] == k {
            return 1
        }
        return 0
    }

    // len nums >= 2
    sums := make([]int, len(nums))
    sum := 0
    result := 0
    for i, num := range nums {
        sum += num
        sums[i] = sum
        if sum == k {
            result += 1
        }
    }

    size := len(nums)

    for j := size-1; j >= 1; j -- {
        for i := j-1 ; i >= 0; i-- {
            if sums[j] - sums[i] == k {
                result += 1
            }
        }
    }

    return result
}
