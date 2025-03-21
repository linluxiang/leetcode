import "slices"

func hIndex(citations []int) int {
    slices.Sort(citations)
    hidx := 0
    for i, v := range citations {
        if i > 0 && v == citations[i-1] {
            continue
        }
        check := min(len(citations) - i, v)
        if check >= hidx {
            hidx = check
        }
    }
    return hidx 
}
