import "slices"

func subsetsWithDup(nums []int) [][]int {
    slices.Sort(nums)

    result := [][]int{}
    path := []int{}
    var dfs func(i int)
    dfs = func(pos int) {
        result = append(result, slices.Clone(path))

        for i := pos; i < len(nums); i++ {
            if i != pos && nums[i] == nums[i-1] {
                // 这里i从pos开始的原因是为了把当前元素添加到path里面
                // 当这个循环结束的时候，i的值是
                continue
            }
            fmt.Println("i: ", i, " pos: ", pos, " path: ", path, " nums[i] ", nums[i])
            path = append(path, nums[i])
            dfs(i+1)
            path = path[:len(path) - 1]
        }
    }

    dfs(0)
    return result
}
