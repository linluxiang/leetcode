import "slices"

func combinationSum(candidates []int, target int) [][]int {
    ans := [][]int{}
    slices.Sort(candidates)

    var dfs func (nums []int, t int, collect []int)

    dfs = func (nums []int, t int, collect []int) {
        if t == 0 {
            ans = append(ans, slices.Clone(collect))
            return
        }

        for _, num := range nums {
            if num > t || (len(collect) > 0 && num > collect[len(collect) - 1]){
                return
            }
            collect = append(collect, num)
            dfs(nums, t - num, collect)
            collect = collect[:len(collect) - 1]
        }
    }


    dfs(candidates, target, []int{})

    return ans
}
