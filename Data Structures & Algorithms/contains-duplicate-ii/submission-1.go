func containsNearbyDuplicate(nums []int, k int) bool {
	rmap := make(map[int]int)

	for i, num := range nums {
		if j, ok := rmap[num]; ok {
			if i - j <= k {
				return true
			}
		}
		rmap[num] = i
	}
	return false
}
