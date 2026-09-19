func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high - low)/2

	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearch(nums, target, low, mid-1)
	}
	return binarySearch(nums, target, mid+1, high)
}
