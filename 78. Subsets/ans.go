import "math"
import "strconv"

func subsets(nums []int) [][]int {
    result := [][]int{}
    cnt := int64(math.Pow(2, float64(len(nums))))
    for i := range cnt {
        r := []int{}
        binary := strconv.FormatInt(i, 2)
        for i:=0; i < len(binary); i++ {
            if binary[i] == '1' {
                r = append(r, nums[len(binary)-i-1])
            }
        }
        result = append(result, r)
    }
    return result
}
