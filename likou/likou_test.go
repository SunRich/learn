package likou

import (
	"fmt"
	"testing"
)

func TestTwoADD(t *testing.T) {
	fmt.Println(twoSum([]int{3, 4, 5, 6}, 11))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("pwwkdfggasdceweaswwfg"))
}

func TestGetMinSlice(t *testing.T) {
	fmt.Println(bsearch([]int{1, 2, 3, 8, 32}, 2))
}

func bsearch(array []int, findVal int) int {
	var arraylen, middleIndex int
	arraylen = len(array)
	if arraylen == 0 {
		return -1
	}

	middleIndex = arraylen / 2
	if array[middleIndex] > findVal {
		return bsearch(array[0:middleIndex], findVal)
	}
	if array[middleIndex] < findVal {
		return bsearch(array[middleIndex:arraylen], findVal)
	}
	if array[middleIndex] == findVal {
		return middleIndex
	}
	return -1
}

func TestFindKthLargest(t *testing.T) {
	result := findKthLargest([]int{7,2, 3, 1, 4, 6}, 2)
	fmt.Println(result)
}
