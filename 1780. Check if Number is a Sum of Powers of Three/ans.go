func checkPowersOfThree(n int) bool {
    mod := 0 
    quotien := 0
    for {
        quotien = n / 3
        mod = n % 3
        n = quotien
        if mod == 2 {
            return false
        }
        if quotien == 0 {
            break
        }
    }
    return true
}
