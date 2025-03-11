import "slices"
import "fmt"

func hasInterval(a, b []int) bool {
    if a[1] > b[0] {
        return true
    }
    return false
}

func minMeetingRooms(intervals [][]int) int {
    if len(intervals) <= 1 {
        return 1
    }

    slices.SortFunc(intervals, func (a, b []int) int {
        return a[0] - b[0]
    })
    
    rooms := [][][]int{
        [][]int{intervals[0]},
    }
    outer:
    for i := 1; i < len(intervals); i++ {
        // fmt.Println("rooms: ", rooms, intervals[i])
        for j, room := range rooms {
            last := room[len(room) - 1]
            if !hasInterval(last, intervals[i]) {
                rooms[j] = append(rooms[j], intervals[i])
                // fmt.Println("append rooms: ", rooms, intervals[i])

                continue outer
            }
        }
        rooms = append(rooms, [][]int{intervals[i]})
    }

    return len(rooms)
}
