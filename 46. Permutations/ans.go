import "slices"

func permute(nums []int) [][]int {
    ans := [][]int{}

    var dfs func(idx int, nums []int, collected []int, visited map[int]bool)

    dfs = func(idx int, nums []int, collected []int, visited map[int]bool) {
        if len(collected) == len(nums) {
            ans = append(ans, slices.Clone(collected))
            return
        }

        for i := 0; i < len(nums); i++ {
            if visited[nums[i]] {
                continue
            }
            collected = append(collected, nums[i])
            visited[nums[i]] = true
            dfs(i+1, nums, collected, visited)
            visited[nums[i]] = false
            collected = collected[:len(collected)-1]
        }
    }

    dfs(0, nums, []int{}, map[int]bool{})

    return ans
}
