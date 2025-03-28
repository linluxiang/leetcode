// import "slices"

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
    // slots := append(slices.Clone(slots1), slots2...)
    // slices.SortFunc(slots, func(a, b []int) int {
    //     if a[0] != b[0] {
    //         return a[0] - b[0]
    //     }
    //     return a[1] - b[1]
    // })

    slices.SortFunc(slots1, func(a, b []int) int {
        return a[0] - b[0]
    })

    slices.SortFunc(slots2, func(a, b []int) int {
        return a[0] - b[0]
    })

    /*
    [a1, a2], [b1, b2], [b3, b4]

    merge a1a2 and b1b2

    1. there is interval, merge a
    2. there is no interval, could a1a2 has intervals with b3b4?, no, so we should remove a1a2
    */

    if slots1[0][0] > slots2[0][0] {
        slots1, slots2 = slots2, slots1
    }

    // the first item of slots1 is smaller than the slots2
    i, j := 0, 0 
    for i < len(slots1)  && j < len(slots2) {
            slot1 := slots1[i]
            slot2 := slots2[j]

            interval := findInterval(slot1, slot2)

            // fmt.Println(slot1, slot2, interval)
            if len(interval) != 0 {
                if interval[1] - interval[0] >= duration {
                    return []int{interval[0], interval[0] + duration}
                }
            }

            // 这个地方非常关键，slot1和slot2中，后面更大的那个还有可能和其他的slot有交叉，但是更小的那个不可能了。
            // 所以更小的这个直接就可以掠过
            if slot1[1] < slot2[1] {
                i++
            } else {
                j++
            }
        }


    return []int{}
}
