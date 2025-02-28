import "fmt"
import "math/rand"

func quickSelect(nums []int, k int) int {
    left := make([]int, 0, len(nums))
    mid := make([]int, 0, len(nums))
    right := make([]int, 0, len(nums))

    pivotIdx := rand.Intn(len(nums))
    pivot := nums[pivotIdx]
    for _, num := range nums {
        if num > pivot {
            left = append(left, num)
        } else if num < pivot {
            right = append(right, num)
        } else {
            mid = append(mid, num)
        }
    }

    if k <= len(left) {
        return quickSelect(left, k)
    }

    if len(left) + len(mid) < k {
        return quickSelect(right, k - len(left) - len(mid))
    }

    return pivot
}

func findKthLargest(nums []int, k int) int {
    return quickSelect(nums, k)
}
