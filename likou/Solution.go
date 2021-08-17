package likou

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		privot := left + (right-left)>>1
		if nums[privot] == target {
			return privot
		} else if nums[privot] > target {
			right = privot - 1
		} else if nums[privot] < target {
			left = privot + 1
		}
	}
	return -1
}
