func count2and5(n int) [2]int {
    result := [2]int{0, 0}
    for n % 2 == 0 {
        n = n / 2
        result[0] += 1
    }

    for n % 5 == 0 {
        n = n / 5
        result[1] += 1
    }

    return result
}

func trailingZeroes(n int) int {
    result := [2]int{0, 0}
    for i := 1; i <= n; i++ {
        r := count2and5(i)
        result[0] += r[0]
        result[1] += r[1]
    }

    if result[0] < result[1] {
        return result[0]
    }
    return result[1]
}
