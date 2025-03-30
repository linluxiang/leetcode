func uniqueOccurrences(arr []int) bool {
    occr := map[int]int{}
    for _, v := range arr {
        occr[v] += 1
    }

    values := map[int]bool{}
    for _, v := range occr {
        if values[v] == true {
            return false
        }
        values[v] = true
    }

    return true
}
