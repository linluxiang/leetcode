import "fmt"

type Heap struct {
	l []int
}

func (h *Heap) Push(v int) {
	h.l = append(h.l, v)
	h.up(len(h.l) - 1)
	// fmt.Println(h.l)
}

func (h *Heap) Pop() int {
	v := h.l[0]
	h.swap(0, len(h.l)-1)
	h.l = h.l[:len(h.l)-1]
	// fmt.Println(h.l)
	h.down(0)
	// fmt.Println(h.l)

	return v
}

func (h *Heap) swap(i, j int) {
	h.l[i], h.l[j] = h.l[j], h.l[i]
}

func (h *Heap) up(idx int) {
	if idx == 0 {
		return
	}
	parent := (idx - 1) >> 1
	if h.l[parent] > h.l[idx] {
		return
	}

	h.swap(idx, parent)
	h.up(parent)
}

func (h *Heap) down(idx int) {
	left := idx*2 + 1
	right := idx*2 + 2
	if left >= len(h.l) {
		// already a leaf node
		return
	}
	swapTarget := left
	// left < len(h.l), has left child
	if right >= len(h.l) {
		swapTarget = left
	} else {
		// has both child
		if h.l[left] > h.l[right] {
			swapTarget = left
		} else {
			swapTarget = right
		}
	}
	if h.l[idx] > h.l[swapTarget] {
		// already a heap, no need to swap
		return
	}
	h.swap(idx, swapTarget)
	h.down(swapTarget)

}

func findKthLargest(nums []int, k int) int {
	heap := &Heap{l: []int{}}
	for _, num := range nums {
		heap.Push(num)
	}

	for range k - 1 {
		heap.Pop()
	}

	return heap.Pop()
}
