import "container/heap"
import "slices"

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	size := len(old)
	last := old[size-1]
	*h = old[:size-1]
	return last
}

func (h *MinHeap) Top() int {
	return (*h)[0]
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	size := len(old)
	last := old[size-1]
	*h = old[:size-1]
	return last
}

func (h *MaxHeap) Top() int {
	return (*h)[0]
}

// 这道题最简单的办法就是每一个window排序。。这样最坏是nk，最好是n*klogk
// 最佳的办法是用一个OrderedList，每次插入删除以后都能保证是有序的。
// 如果用红黑树实现，那么每次删除和插入稳定都是logk，总体就是nlogk
// 如果是用列表实现，那么最坏的情况是每次插入到头部，就要移动k，还是nk
// 现在的思路是用两个堆，一个最小堆和一个最大堆，最大堆用来存左半边，最小堆用来存右半边

func medianSlidingWindow(nums []int, k int) []float64 {
	ans := []float64{}
    if k == 1 {
        for _, num := range nums {
        ans = append(ans, float64(num))

        }
        return ans    
    }

	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}

	var heapSize int
    heapSize = k / 2

	firstWindow := []int{}

	for i := range k {
		firstWindow = append(firstWindow, nums[i])
	}

	slices.Sort(firstWindow)
    // 根据第一个窗口建立左堆和右堆

	for i := range firstWindow[:heapSize]{
		heap.Push(maxHeap, firstWindow[i])
	}

    for i := heapSize; i < len(firstWindow); i++ {
        heap.Push(minHeap, firstWindow[i])
    }

	for i := 0; i+k <= len(nums); i++ {
		if i > 0 {
			toRemove := nums[i-1]
			toAdd := nums[i+k-1]

			idx := slices.Index(*minHeap, toRemove)
			if idx != -1 {
				heap.Remove(minHeap, idx)
			} else {
				idx = slices.Index(*maxHeap, toRemove)
				heap.Remove(maxHeap, idx)
			}

            // 如果k等于2和3，那么有可能其中一个是空
			if maxHeap.Len() == 0 {
                if toAdd < minHeap.Top() {
                    heap.Push(maxHeap, toAdd)
                } else {
				    heap.Push(minHeap, toAdd)
                }
			} else if minHeap.Len() == 0 {
                if toAdd > maxHeap.Top() {
				    heap.Push(minHeap, toAdd)
                } else {
				    heap.Push(maxHeap, toAdd)
                }
			} else {
				if toAdd > maxHeap.Top() {
					heap.Push(minHeap, toAdd)
				} else {
					heap.Push(maxHeap, toAdd)
				}
			}
            // 这里需要平衡两个堆
			if minHeap.Len()-maxHeap.Len() > 1 {
				v := heap.Pop(minHeap).(int)
				heap.Push(maxHeap, v)
			} else if maxHeap.Len()-minHeap.Len() > 1 {
				v := heap.Pop(maxHeap).(int)
				heap.Push(minHeap, v)
			}

		}

        if maxHeap.Len() == minHeap.Len() {
		    median := float64(minHeap.Top()+maxHeap.Top()) / 2
		    ans = append(ans, median)
        } else {
            // 这里的技巧是如果两个堆不一样，那么就是取多的那个
            if maxHeap.Len() > minHeap.Len() {
                ans = append(ans, float64(maxHeap.Top()))
            } else {
                ans = append(ans, float64(minHeap.Top()))
            }
        }

	}

	return ans
}
