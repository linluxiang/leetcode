import "container/heap"

type Pair struct {
    X,Y int
    Sum int
}

type Heap struct {
    pairs []*Pair
    k int
}

func NewHeap(k int) *Heap {
    return &Heap{
        pairs: []*Pair{},
        k: k,
    }
}

func (h *Heap) Pairs() []Pair {
    r := []Pair{}
    for _, p := range h.pairs {
        r = append(r, *p)
    }
    return r
}

func (h *Heap) Len() int {
    return len(h.pairs)
}

func (h *Heap) Less(i, j int) bool {
    return h.pairs[i].Sum <= h.pairs[j].Sum
}

func (h *Heap) Swap(i, j int) {
    h.pairs[j], h.pairs[i] = h.pairs[i], h.pairs[j]
}

func (h *Heap) Push(x any) {
    h.pairs = append(h.pairs, x.(*Pair))
}

func (h *Heap) Pop() any {
    if len(h.pairs) == 0 {
        return nil
    }
    // the swap has already been called here, so return the last one
    result := h.pairs[len(h.pairs) - 1]
    h.pairs = h.pairs[0:len(h.pairs) -1]
    return result
}

/*
本质上可以把整个表看成多个有序数列的归并排序。
常规的归并排序算法只有两个队列，做法是比较a队列的头和b队列的头，哪个小的就放进结果队列。
而这个题目可以看成是如下的n个队列做归并
a0b0, a0b1, ..., a0bm, m == N2-1
a1b0, a1b1, ..., a1bm, m == N2-1
...
anb0, anb1, ..., anbm, n == N1-1

这个时候用最小堆可以方便的比较n个头
所以第一步就是先把n个头加入最小堆，然后弹出堆头以后，那个aibj把他的右边加入到最小堆里面进行下一轮比较
可以优化的点是第一步不需要把所有的aib0都加入堆，因为他是递增的，所以直到某一行被弹出以后再加。
不过为了思路简单，可以直接把所有的aib0都加入堆
*/

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    h := NewHeap(k)
    heap.Init(h)
    // count := 0
    result := [][]int{}
    for i := range nums1 {
        heap.Push(h, &Pair{i, 0, nums1[i]+nums2[0]})
    }

    for range k {
        v := heap.Pop(h)

        pair := v.(*Pair)
        result = append(result, []int{nums1[pair.X], nums2[pair.Y]})
        if pair.Y + 1 > len(nums2) - 1 {
            // 该队列已经合并完了
            continue
        }
        heap.Push(h, &Pair{pair.X, pair.Y+1, nums1[pair.X] + nums2[pair.Y+1]})
    }
    return result
}
