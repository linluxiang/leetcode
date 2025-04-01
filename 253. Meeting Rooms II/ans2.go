import "slices"
func minMeetingRooms(intervals [][]int) int {
    slices.SortFunc(intervals, func (a, b []int) int {
        if a[0] != b[0] {
            return a[0] - b[0]
        }

        return a[1] - b[1]
    })

    meetingRooms := [][]int{}
loop:
    for _, entry := range intervals {
        for i, m := range meetingRooms {
            if m[1] <= entry[0] {
                // can use the same room
                meetingRooms[i] = entry
                continue loop
            }
        }
        meetingRooms = append(meetingRooms, entry)
    }

    return len(meetingRooms)
}
