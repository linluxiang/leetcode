func swap(num []int, i, j int) {
	temp := num[i]
	num[i] = num[j]
	num[j] = temp
}

var magic = -99999

func findFirstInvalid(nums []int, i int) (int, bool) {
	if i == len(nums)-1 {
		return -1, false
	}
	for i < len(nums) {
		if nums[i] == magic {
			break
		}
		i++
	}
	if i >= len(nums) {
		return -1, false
	}
	return i, true
}

func findFirstValidAfter(nums []int, i, j int) (int, bool) {
	if i == len(nums)-1 {
		return -1, false
	}
	if j <= i {
		j = i + 1
	}
	for j < len(nums) {
		if nums[j] != magic {
			break
		}
		j++
	}
	if j >= len(nums) {
		return -1, false
	}
	return j, true
}

func removeDuplicates(nums []int) int {
	existed := make(map[int]byte)
	amount := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		count, ok := existed[v]
		if !ok {
			existed[v] = 1
			amount += 1
			continue
		}
		if count == 1 {
			existed[v] = 2
			amount += 1
			continue
		}

		nums[i] = magic
	}

	var i, j = 0, 0
	var ok bool
	for i < len(nums)-1 {
		i, ok = findFirstInvalid(nums, i)
		if !ok {
			break
		}
		j, ok = findFirstValidAfter(nums, i, j)
		if !ok {
			break
		}
		swap(nums, i, j)
	}

	return amount
}
