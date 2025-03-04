func subarraySum(nums []int, k int) int {
    if len(nums) == 1 {
        if nums[0] == k {
            return 1
        }
        return 0
    }

    // len nums >= 2
    // sums[i, j] = sums[j] - sums[i-1]
    // 如果要表示从0开始的子数组的和，等于当前坐标的前缀和减去 (-1)位置的前缀和。。-1位置的前缀和一定是0，所以cache[0]一定至少出现1次

    sums := make([]int, len(nums) )
    sum := 0
    result := 0
    
    // Cache表示的是某个前缀和出现的次数
    cache := map[int]int{}

    cache[0] = 1 // -1位置的前缀和一定是0，那么0至少是出现1次

    for i, num := range nums {
        sum += num
        sums[i] = sum
        // cache[sum] += 1 不能预先计算某个sum出现次数，理由在下面说
    }

    // fmt.Println("sums: ", sums, ". cache: ", cache)
    for _, v := range sums {
        //计算在当前位置以前，v-k有没有出现，如果有就加上，当前位置以前很重要，所以不能预先算cache，必须走一步把当前的添加进去。
        if cnt, ok := cache[v-k]; ok {
            result += cnt
            // fmt.Println("exist: ", v, v-k, result)
        }
        //
        cache[v] += 1 // 把当前加进去方便后面的计算。
    }

    return result
}
