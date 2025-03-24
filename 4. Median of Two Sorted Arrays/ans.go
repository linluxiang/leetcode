// import "slices"

func binarySearch(nums []int, target int) (idx int) {
    i, j := -1, len(nums)
    for i < j - 1 {
        mid := i + (j-i) >> 1
        if nums[mid] <= target {
            i =  mid
        } else {
            j = mid
        }
    }
    return j
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	if len(nums1) == 0 || len(nums2) == 0 {
		var nonEmpty []int
		if len(nums1) == 0 {
			nonEmpty = nums2
		}

		if len(nums2) == 0 {
			nonEmpty = nums1
		}

		if len(nonEmpty)%2 == 1 {
			return float64(nonEmpty[len(nonEmpty)/2])
		} else {
			return float64(nonEmpty[len(nonEmpty)/2-1]+nonEmpty[len(nonEmpty)/2]) / 2
		}
	}

	if nums1[0] > nums2[0] {
		nums1, nums2 = nums2, nums1
	}

	size1 := len(nums1)
	size2 := len(nums2)
	totalSize := size1 + size2
	isOdd := totalSize%2 == 1
	var medianIdx int
	if isOdd {
		medianIdx = totalSize / 2
	} else {
		medianIdx = totalSize / 2 - 1
	}

	for medianIdx != 0 {
		idx := binarySearch(nums1, nums2[0])

        if idx == len(nums1) {
            if medianIdx < len(nums1) {
                nums1 = nums1[medianIdx:]
                medianIdx = 0
            } else { // medianIdx >= len(nums1)
                medianIdx -= len(nums1)
                nums1 = nums1[0:0]
            }
        } else {
            if medianIdx > idx {
                medianIdx -= idx
		        nums1 = nums1[idx:len(nums1)]
            } else {
                nums1 = nums1[medianIdx:]
                medianIdx = 0
            }

        }


		if len(nums1) == 0 || len(nums2) == 0 {
			break
		}

		if nums1[0] > nums2[0] {
			nums1, nums2 = nums2, nums1
		}
	}

	if len(nums1) == 0 || len(nums2) == 0 {
		var nums []int
		if len(nums1) == 0 {
			nums = nums2
		} else {
			nums = nums1
		}
		if isOdd {
			return float64(nums[medianIdx])
		} else {
			return float64(nums[medianIdx]+nums[medianIdx+1]) / 2.0
		}
	}

	// nums1 and nums2 are both not equal to 0

	// medianIdx == 0
	if isOdd {
		return float64(nums1[0])
	} else {
		if len(nums1) == 1 {
			return float64(nums1[0]+nums2[0]) / 2.0
		} else {
			return float64(nums1[0]+min(nums1[1], nums2[0])) / 2.0
		}
	}
}
