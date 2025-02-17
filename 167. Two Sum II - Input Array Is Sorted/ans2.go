func twoSum(numbers []int, target int) []int {
    i := 0
    j := len(numbers) - 1

    for {
        sum := numbers[i] + numbers[j]
        if sum == target {
            return []int{i+1, j+1}
        } else if sum < target {
            i += 1
        } else if sum > target {
            j -= 1
        }
    }
    panic("should not happen")
}

