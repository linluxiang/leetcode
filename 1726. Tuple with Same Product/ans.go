import "fmt"
// import "slices"


func tupleSameProduct(nums []int) int {
    if len(nums) < 4 {
        return 0
    }

    prod := map[int]int{}
    cnt := 0
    size := len(nums)
    for i := 0; i < size - 1; i++ {
        for j := i+1; j < size; j++ {
            v := nums[i] * nums[j]
            if _, ok := prod[v]; ok {
                prod[v] += 1
            } else {
                prod[v] = 1
            }
        }
    }

    for _, v := range prod {
        if v == 1 {
            continue
        }
        cnt += v * (v - 1) / 2

    }
    return cnt * 8
}
