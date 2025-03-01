// import "container/heap"

func power(x int, p int) int {

    if x == 1 {
        return p
    }

    if x % 2 == 1 {
        return power(x * 3 + 1, p+1)
    } else {
        return power(x / 2, p+1)
    }
}

func quickSelect(nums [][]int, k int) []int {
    left := make([][]int, 0, len(nums))
    mid := make([][]int, 0, len(nums))
    right := make([][]int, 0, len(nums))   

    idx := rand.Intn(len(nums))
    pivot := nums[idx]

    for _, num := range nums {
        if num[1] < pivot[1] {
            left = append(left, num)
        }

        if num[1] == pivot[1] {
            if num[0] < pivot[0] {
                left = append(left, num)
            } else if num[0] > pivot[0] {
                right = append(right, num)
            } else {
                mid = append(mid, num)
            }
        }

        if num[1] > pivot[1] {
            right = append(right, num)
        }
    }

    if len(left) >= k {
        return quickSelect(left, k)
    }

    if len(left) + len(mid) < k {
        return quickSelect(right, k - len(left) - len(mid))
    }

    // k is in mid
    return pivot
} 

// type Heap [][]int

// func (h Heap) Len() int {
//     return len(h)
// }

// func (h Heap) Swap(i, j int) {
//     h[i], h[j] = h[j], h[i]
// }

// func (h Heap) Less(i, j int) bool {
//     if h[i][1] < h[j][1] {
//         return true
//     }
//     if h[i][1] == h[j][1] {
//         return h[i][0] < h[j][0]
//     }
//     return false
// }

// func (h *Heap) Push(v any) {
//     *h = append(*h, v.([]int))
// }

// func (h *Heap) Pop() any {
//     old := *h
//     n := len(old)
//     last := old[n - 1]
//     *h = old[:n-1]
//     return last
// }

func getKth(lo int, hi int, k int) int {
    nums := make([][]int, 0, hi - lo + 1)

    // h := &Heap{}
    // heap.Init(h)
    for i := lo; i <= hi; i++ {
        p := power(i, 0)
        nums = append(nums, []int{i, p})
        // heap.Push(h, []int{i, p})
    }
    // fmt.Println(*h)
    return quickSelect(nums, k)[0]
    // var result int
    // for range k {
    //     v := heap.Pop(h)
    //     vv := v.([]int)
    //     result = vv[0]
    // }

    // return result
}
