type Stack []int

func (s *Stack) Top() int {
	if len(*s) == 0 {
		return -1
	}

	return (*s)[len(*s)-1]
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
    // fmt.Println("after push: ", s)
}

func (s *Stack) Pop() int {
    old := *s
	if len(old) > 0 {
		v := old[len(old)-1]
		*s = old[:len(old)-1]
		return v
	}
	return -1
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	s := &Stack{}
	nextGreater := map[int]int{}

	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
        // fmt.Println(i, num, s)

		for len(*s) > 0 && s.Top() < num {
			s.Pop()
		}

		if len(*s) == 0 {
			nextGreater[num] = -1
			s.Push(num)
			continue
		}

		nextGreater[num] = s.Top()
		s.Push(num)
	}

	ans := make([]int, len(nums1))
	for i := range nums1 {
		ans[i] = nextGreater[nums1[i]]
	}

	return ans
}
