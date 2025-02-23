import "fmt"

func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    if len(nums) == 1 {
        return nums[0]
    }

    currSubarray := nums[0]
    result := nums[0]
    for i:=1; i < len(nums); i++ {
        if currSubarray < 0 {
            currSubarray = nums[i]
        } else {
            currSubarray += nums[i]
        }

        if currSubarray > result {
            result = currSubarray
        }

    }
    return result
}
