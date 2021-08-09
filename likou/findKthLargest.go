package likou

func findKthLargest(nums []int, k int) int {

}

func quickSort(nums []int, left, right, k int) int {
	partition := partition(nums, left, right)
	if len(nums)-1-k == partition {
		return partition
	}
	if len(nums)-1-k < partition {
		return quickSort(nums,left,partition-1)
	}

}

func partition(nums []int, left, right int) int {

}
