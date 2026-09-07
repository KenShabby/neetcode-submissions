func majorityElement(nums []int) int {
	counter := make(map[int]int)
	for _, curr := range nums{
		counter[curr]++
	}

	max := nums[0]
	for num, count := range counter {
		if count > counter[max] {
			max = num
		}
	}

	return max
}
