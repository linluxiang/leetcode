import "slices"

func subsetsWithDup(nums []int) [][]int {
    slices.Sort(nums)

    result := [][]int{[]int{}}
    path := []int{}
    var dfs func(i int)

    dfs = func(pos int) {

        path = append(path, nums[pos])

        for i := pos+1; i < len(nums); i++ {
            if i>pos+1 && nums[i] == nums[i-1] {
                // 这里i必须大于pos + 1的原因是，如果第一次出现一个数，我们可以继续处理，只有当第二次出现的时候，我们才忽略他。
                continue
            }

            dfs(i)
            path = path[:len(path) - 1]

        }
        result = append(result, slices.Clone(path))

    }

    for i := range nums {
        if i >= 1 && nums[i] == nums[i-1] {
            continue
        }
        path = []int{}
        dfs(i)
    }

    return result
}
