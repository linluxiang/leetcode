func search(numbers []int, target, i, j int) (index int) {
    if i == j - 1 {
        if numbers[i] == target {
            return i
        } else {
            return -1
        }
    }

    mid := i + (j - i) >> 1
    
    if target < numbers[mid] {
        return search(numbers, target, i, mid)
    } else {
        return search(numbers, target, mid, j)
    }

}

func twoSum(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        j := search(numbers, target - numbers[i], i, len(numbers))
        if j != -1 {
            return []int{i+1, j+1}
        }
    }
    panic("should not happen")
}
