import "fmt"

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }

    if n == 2 {
        return 2
    }

    steps := make([]int, n+1)

    steps[0] = 0
    steps[1] = 1
    steps[2] = 2
    for i := 3; i <= n; i++ {
        steps[i] = steps[i-2] + steps[i-1]
    }
    // fmt.Println(steps)
    return steps[n]
}
