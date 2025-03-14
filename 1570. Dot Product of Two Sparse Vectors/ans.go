type Set map[int]int

func (s Set) Intersect(b Set) []int {
    var smaller,bigger Set
    if len(s) < len(b) {
        smaller, bigger = s, b
    } else {
        smaller, bigger = b, s
    }
    r := make([]int, 0, len(s))
    for k := range smaller {
        if _, ok := bigger[k]; ok{
            r = append(r, k)
        }
    }

    return r
}


type SparseVector struct {
    set Set
}

func Constructor(nums []int) SparseVector {
    v := map[int]int{}
    for idx, num := range nums {
        if num == 0 {
            continue
        }
        v[idx] = num
    }

    return SparseVector{
        set: Set(v),
    }
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    set := this.set.Intersect(vec.set)
    product := 0
    for _, idx := range set {
        product += this.set[idx] * vec.set[idx]
    }
    return product
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
