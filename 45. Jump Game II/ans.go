func jump(nums []int) int {
    MAX := 100000
    minJump := make([]int, len(nums))
    for i := range nums {
        minJump[i] = MAX
    }
    minJump[0] = 0

    for i, v := range nums {
        for j := i + 1; j <= i + v && j < len(nums); j++ {
            if minJump[j] > minJump[i] + 1 {
                minJump[j] = minJump[i] + 1
            }
        }
        // fmt.Println(minJump)
    }

    return minJump[len(minJump) - 1]
}
