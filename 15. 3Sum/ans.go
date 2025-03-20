import "slices"
func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    ans := [][]int{}
    for i := range nums {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j := i + 1
        k := len(nums) - 1
        for j < k {
            if nums[i] + nums[j] + nums[k] < 0 {
                j++
                continue
            }
            if nums[i] + nums[j] + nums[k] > 0 {
                k--
                continue
            }
            ans = append(ans, []int{nums[i], nums[j], nums[k]})
            j++
            k--
            for j < k && nums[j] == nums[j-1] {
                j++
            }
            for j < k && nums[k] == nums[k+1] {
                k--
            }
            // fmt.Println("after move: ", i, j, k, ans)

        }
    }
    return ans
}
