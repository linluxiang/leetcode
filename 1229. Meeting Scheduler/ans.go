import "slices"

func findInterval(slot1 []int, slot2[]int) []int {
    // make sure that slot1 starts earlier
    if slot1[0] > slot2[0] {
        slot1, slot2 = slot2, slot1
    }

    if slot1[1] > slot2[1] {
        return slot2
    }

    if slot1[1] > slot2[0] {
        return []int{slot2[0], slot1[1]}
    }

    // there is no interval
    return []int{}
}

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
    // slots := append(slots1.clone(), slots2...)

    slices.SortFunc(slots1, func(a, b []int) int {
        return a[0] - b[0]
    })

    slices.SortFunc(slots2, func(a, b []int) int {
        return a[0] - b[0]
    })

    intervals := [][]int{}
    /*
    [a1, a2], [b1, b2], [b3, b4]

    merge ab and cd

    1. there is interval, result to merge e, f
    2. there is no interval, then we will try to merge cd and ef
    */
    // for _, slot := range slots {

    // }

    if slots1[0][0] > slots2[0][0] {
        slots1, slots2 = slots2, slots1
    }

    // the first item of slots1 is smaller than the slots2

    for i := range slots1 {
        for j := range slots2 {
            slot1 := slots1[i]
            slot2 := slots2[j]
            if slot1[0] > slot2[1] {
                continue
            }

            if slot2[0] > slot1[1] {
                break
            }

            interval := findInterval(slot1, slot2)

            // fmt.Println(slot1, slot2, interval)
            if len(interval) != 0 {
                intervals = append(intervals, interval)
            }
        }
    }
    // fmt.Println("intervals: ", intervals)
    for _, interval := range intervals {
        if interval[1] - interval[0] >= duration {
            return []int{interval[0], interval[0] + duration}
        }
    }

    return []int{}
}
