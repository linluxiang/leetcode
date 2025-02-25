import "fmt"

func search(nums []int, i, j int) int {
    if j - i + 1 <= 3 {
        if nums[i] > nums[i+1] {
            return i
        }
        //nums[i] < nums[j]
        if nums[i+1] > nums[j] {
            return i+1
        }

        return j
    }

    mid := i+(j-i) >>1
    // fmt.Printf("%d %d\n", i, j)
    if nums[mid - 1] > nums[mid] {
        return search(nums, i, mid)
    }
    // nums[mid - 1] < nums[mid]
    if nums[mid] > nums[mid + 1] {
        return mid
    }

    return search(nums, mid, j)
}

func findPeakElement(nums []int) int {
    if len(nums) == 1 {
        return 0
    }

    if len(nums) == 2 {
        if nums[0] > nums[1] {
            return 0
        } else {
            return 1
        }
    }

    return search(nums, 0, len(nums) - 1)

}
