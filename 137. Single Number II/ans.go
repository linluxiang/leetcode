import "fmt"

func singleNumber(nums []int) int {
    /*
    整体思路是这样，统计所有的数每一个二进制bit上面出现了多少次1。
    统计完以后对于所有的出现次数mod 3处理。那么出现余数的肯定就是多出来的那个数字
    因为他只会出现1次，所以mod 3以后就是1，那么就可以以这一位为1进行二进制计算来还原出多余的那个数
    */
    appears := [32]int32{} // 表示每个bit上出现1的次数
    for _, num := range nums {
        n := int32(num)
        for i := 0; i < 32; i ++ {
            appears[i] += n & 1
            n = n >> 1
        }
    }

    // fmt.Println(appears)
    for i := range appears {
        appears[i] = appears[i] % 3
    }
    
    // fmt.Println(appears)

    result := int32(0)
    for i, v := range appears {
        result |= int32(v) << i
    } 
    return int(result)
}
