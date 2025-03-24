func binarySearchFirst(nums []int, target int) int {
    i, j := -1, len(nums)
    for i < j - 1 {
        mid := i + (j-i) >> 1

        if nums[mid] < target {
            i = mid
        } else if nums[mid] >= target {
            j = mid
        }
    }
    return j
}

func binarySearchLast(nums []int, target int) int {
    i, j := -1, len(nums)
    for i < j -1 {
        mid := i + (j-i) >> 1

        if nums[mid] <= target {
            i = mid
        } else if nums[mid] > target {
            j = mid
        }
    }
    return i
}

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    first := binarySearchFirst(nums, target)
//    fmt.Println("first: ", first)
    if first >= len(nums) || nums[first] != target {
        first = -1
    }

    last := binarySearchLast(nums, target)
 //   fmt.Println("last: ", last)
    if last < 0 || nums[last] != target {
        last = -1
    }

    return []int{first, last}

}
