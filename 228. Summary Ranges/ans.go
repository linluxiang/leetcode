func summaryRanges(nums []int) []string {
    if len(nums) == 1 {
        return []string{fmt.Sprintf("%d", nums[0])}
    }
    ans := []string{}
    start := 0
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i+1] != nums[i] + 1 {
            if start == i {
                ans = append(ans, fmt.Sprintf("%d", nums[start]))
            } else {
                ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[i]))
            }
            start = i+1
            continue
        }
    }

    if start == len(nums) - 1 {
        ans = append(ans, fmt.Sprintf("%d", nums[start]))
    } else if start < len(nums) - 1 {
        ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[len(nums) - 1]))
    }

    return ans
}
