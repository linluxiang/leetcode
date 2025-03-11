func binarySearch(nums []int, target, i, j int) int {
	if nums[i] == target {
		return i
	}
	if nums[j] == target {
		return j
	}

	mididx := i + (j-i)>>1
	if nums[mididx] == target {
		return mididx
	}
	first, mid, last := nums[i], nums[mididx], nums[j]
	if first <= mid && mid <= last {
		if target < first || target > last {
			return -1
		}
		if target > mid {
			return binarySearch(nums, target, mididx+1, j)
		} else {
			return binarySearch(nums, target, i, mididx-1)
		}
	} else if first <= mid && mid >= last {
		// the biggest is in the middle
		if target < first || target > mid {
			// search right
			return binarySearch(nums, target, mididx+1, j)
		} else {
			// search left
			return binarySearch(nums, target, i, mididx-1)
		}
	} else { // first >= mid && mid <= last
        if target >= first || target < mid {
            //search left
            return binarySearch(nums, target, i, mididx-1)
        } else {
            return binarySearch(nums, target, mididx+1, j)

        }
    }
}

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}
