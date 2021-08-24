package likou

//215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k)
}

func quickSort(nums []int, left, right, k int) int {
	partition := partition(nums, left, right)
	if len(nums)-k == partition {
		return nums[partition]
	}
	if len(nums)-k < partition {
		return quickSort(nums, left, partition-1, k)
	}

	return quickSort(nums, partition+1, right, k)
}

func partition(nums []int, left, right int) int {
	//partition := rand.Int()%(right-left+1) + left
	partitionIndex:=left
	index := partitionIndex + 1
	for i := index; i <= right; i++ {
		if nums[i] < nums[partitionIndex] {
			swap(nums, i, index)
			index++
		}
	}
	swap(nums, partitionIndex, index-1)
	return index - 1
}

func swap(result []int, i, j int) {
	if i == j {
		return
	}
	result[i] = result[i] ^ result[j]
	result[j] = result[i] ^ result[j]
	result[i] = result[i] ^ result[j]
}
