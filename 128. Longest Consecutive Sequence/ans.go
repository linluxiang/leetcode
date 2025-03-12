func findLongest(n int, longest map[int]int, existed map[int]bool) {
    if v, ok := longest[n-1]; ok {
        longest[n] = v + 1
        return
    }
    if existed[n-1] == false {
        longest[n] = 1
        return
    }
    findLongest(n-1, longest, existed)
    longest[n] = longest[n-1] + 1
}

func longestConsecutive(nums []int) int {
    existed := map[int]bool {}
    longest := map[int]int {}
    maxV := 0

    for _, v := range nums {
        existed[v] = true
    }

    for _, v := range nums {
        findLongest(v, longest, existed)
        if longest[v] > maxV {
            maxV = longest[v]
        }
    }
    return maxV
} 
