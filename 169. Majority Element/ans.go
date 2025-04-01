func majorityElement(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    appears := map[int]int{
    }

    for _, num := range nums {
        appears[num] += 1
        if appears[num] > len(nums)/2 {
            return num
        }
    }
    return 0
}
