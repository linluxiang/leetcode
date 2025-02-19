
type Window struct {
	i, j int
	nums []int
	sum  int
    ans int
}

func (w *Window) ExpandRightOnce(target int) bool {
	if w.sum >= target {
		// No need to expand
		return false
	}

	if w.j == len(w.nums)-1 {
		return false
	}

	w.j += 1
	w.sum += w.nums[w.j]
    w.ans = w.j - w.i + 1

	return true
}

// return the minimum distance in this round of Narrow
func (w *Window) NarrowLeftOnce(target int) bool {
	if w.sum < target {
		// No need to narrow down
		return false
	}
	if w.i == w.j {
		return false
	}

	w.sum -= w.nums[w.i]
	w.i += 1
    w.ans = w.j - w.i + 1
	return true
}

func minSubArrayLen(target int, nums []int) int {
	window := &Window{
		i:    0,
		j:    0,
		sum:  nums[0],
		nums: nums,
        ans: 1,
	}
	ans := 999999
	for {
        var expandedOnce, narrowDownedOnce bool
		for {
			if !window.ExpandRightOnce(target) {
				break
			}

            expandedOnce = true
			if window.sum >= target {
				if window.ans < ans {
					ans = window.ans
				}
			}

		}
		for {
			if !window.NarrowLeftOnce(target) {
				break
			}
            narrowDownedOnce = false

			if window.sum >= target {
				if window.ans < ans {
					ans = window.ans
				}
			}
		}

		if !expandedOnce && !narrowDownedOnce {
            if window.sum >= target {
				if window.ans < ans {
					ans = window.ans
				}
			}
            break
		}

	}

	if ans == 999999 {
		return 0
	}
	return ans
}
