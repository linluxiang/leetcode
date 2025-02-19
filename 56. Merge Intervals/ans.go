import "slices"

func merge2(left, right []int) []int {
	if left[0] > right[0] {
		left, right = right, left
	}

	if left[1] < right[0] {
		return []int{}
	}

	if left[1] > right[1] {
		return []int{left[0], left[1]}
	}

	return []int{left[0], right[1]}
}

func search(intervals [][]int, target, i, j int) int {
	// Find the first item whose first entry is bigger or equal than target
	if i == j-1 {
		return j
	}
	mid := i + (j-i)>>1

	if intervals[mid][0] >= target {
		return search(intervals, target, i, mid)
	} else {
		return search(intervals, target, mid, j)
	}
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	ans := [][]int{}
	for _, interval := range intervals {
		if len(ans) == 0 {
			ans = append(ans, interval)
			continue
		}

		idx := search(ans, interval[0], -1, len(ans))
		if idx == len(ans) {
			ans = append(ans, interval)
		} else {
			ans = slices.Insert(ans, idx, interval)

		}

		var i int = idx
		for i + 1 < len(ans) {
			result := merge2(ans[i], ans[i+1])
			if len(result) == 0 {
				break
			} else {
				ans[i] = result
                j := i + 2
                if j > len(ans) {
                    j = len(ans)
                }
                ans = slices.Delete(ans, i+1, j)
			}
		}
        i = idx - 1
		for i >= 0 {
			result := merge2(ans[i], ans[i+1])
            if len(result) == 0 {
                break
            }

		    ans[i] = result
            j := i + 2
                if j > len(ans) {
                    j = len(ans)
                }
			ans = slices.Delete(ans, i+1, j)
            i--
		}
	}

	return ans
}
