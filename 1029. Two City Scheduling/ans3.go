// import "slices"

func twoCitySchedCost(costs [][]int) int {
    mem := map[[2]int]int{}

    var choose func(a, b int) int
    choose = func(a, b int) int {
        // a, b 分别是a个人选了a城市，b个人选了b城市
        if a == 0 && b == 0{
            return 0
        }

        k := [2]int{a, b}
        if v, ok := mem[k]; ok {
            return v
        }
        if a == 0 {
            v := choose(a, b-1) + costs[b-1][1]
            mem[k] = v
            return v
        }

        if b == 0 {
            v := choose(a-1, b) + costs[a-1][0]
            mem[k] = v
            return v
        }

        option1 := choose(a-1, b) + costs[a+b-1][0]
        option2 := choose(a, b-1) + costs[a+b-1][1]
        v := min(option1, option2)
        mem[k] = v
        return v
    }
    
    n := len(costs) / 2
    return choose(n, n)
}
