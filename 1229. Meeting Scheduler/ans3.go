import "slices"
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
    slices.SortFunc(slots1, func (a, b []int) int {
        return a[0] - b[0]
    })

    slices.SortFunc(slots2, func (a, b []int) int {
        return a[0] - b[0]
    })

    i, j := 0, 0

    for i < len(slots1) && j < len(slots2) {
        a := slots1[i]
        b := slots2[j]
        start := max(a[0] ,b[0])
        end := min(a[1], b[1])
        if start < end && end - start >= duration {
            return []int{start, start+duration}
        }

        if a[1] > b[1] {
            j+=1
        } else {
            i+=1
        }
    }

    return []int{}
}
