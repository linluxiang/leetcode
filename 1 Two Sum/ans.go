func twoSum(nums []int, target int) []int {
    ans := make([]int, 0, 2)
    if len(nums) == 0 {
        return ans
    }
    if len(nums) == 2 {
        return append(ans, 0, 1)
    }

    index := make(map[int]int, len(nums))
    for i, v := range nums {
        index[v] = i
    }

    for i, v := range nums {
        if idx, ok := index[target - v]; ok {
            if idx == i {
                continue
            }
            ans = append(ans, i, idx)
            break
        }
    }
    return ans 
}
