import "slices"
func minMeetingRooms(intervals [][]int) int {
    slices.SortFunc(intervals, func (a, b []int) int {
        if a[0] != b[0] {
            return a[0] - b[0]
        }

        return a[1] - b[1]
    })

    meetingRooms := [][]int{} // 这里，会议室可以用一个最小堆来组织，堆的顶端是最早结束的会议，所以堆的比较用的是interval[1]
loop:
    for _, entry := range intervals {
        for i, m := range meetingRooms { // 在这里只需要查看堆顶的元素是否有空，有空的话就弹出。
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
