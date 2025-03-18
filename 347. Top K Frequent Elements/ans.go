import "container/heap"

type Node [2]int

type Heap struct {
    data []Node
    k int
}

func (h Heap) Len() int {
    return len(h.data)
}

func (h Heap) Swap(i, j int) {
    h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h Heap) Less(a, b int) bool {
    return h.data[a][0] < h.data[b][0]
}

func (h *Heap) Push(nodeAny any) {
    node := nodeAny.(Node)
    h.data = append(h.data, node)
    // fmt.Println("data: ", h.data)
}

func (h *Heap) Pop() any {
    if h.Len() == 0 {
        return nil
    }
    size := len(h.data)
    last := h.data[size - 1]
    h.data = h.data[:size - 1]
    return last
}

// func (h *Heap) Debug() {

// }

func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{} // num to freq
    
    for _, num := range nums {
        freq[num] += 1
    }

    // fmt.Println(freq)

    h := &Heap{
        data: []Node{},
        k: k,
    }

    for k, v := range freq {
        node := Node{v, k}
        heap.Push(h, node)
        // Important, the pop mustn't be written in the Heap.Push function
        if len(h.data) > h.k {
            heap.Pop(h)
        }

    }

    result := []int{}

    for range k {
        v := heap.Pop(h)
        node := v.(Node)
        result = append(result, node[1])
    }

    return result
}
