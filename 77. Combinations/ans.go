import "slices"

func combine(n int, k int) [][]int {
    ans := [][]int{}

    collect := []int{}

    var dfs func(start int, n int, k int, collect []int)

    dfs = func(start int, n int, k int, collect []int) {
        // if len(collect) == k {
        //     ans = append(ans, slices.Clone(collect))
        //     return
        // }
        for i := start; i <= n; i++ {
            collect = append(collect, i)
            if len(collect) == k {
                ans = append(ans, slices.Clone(collect))
            } else {
                dfs(i+1, n, k, collect) // collect has already been added to ans
            }
            collect = collect[:len(collect) - 1]
        }
    }

    dfs(1, n, k, collect)
    return ans
}
