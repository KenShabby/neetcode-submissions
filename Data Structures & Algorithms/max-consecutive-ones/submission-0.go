func findMaxConsecutiveOnes(nums []int) int {
	var count, max int

	for i := range nums {
		if nums[i] == 1 {
			count++
			if count > max{
				max = count
			}
		} else {
			if count > max{
				max = count
			}
			count = 0
		}
	}
	return max
}
