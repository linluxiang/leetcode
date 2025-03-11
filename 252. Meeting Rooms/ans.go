import "slices"
func canJoinBoth(a, b []int) bool {
    if a[1] <= b[0] {
        return true
    } 
    return false
}
func canAttendMeetings(intervals [][]int) bool {
    if len(intervals) <= 1 {
        return true
    }

    slices.SortFunc(intervals, func (a, b []int) int {
        return a[0] - b[0]
    })

    for i := 0; i < len(intervals) - 1; i++ {
        if !canJoinBoth(intervals[i], intervals[i+1]) {
            return false
        }
    }
    return true
}
