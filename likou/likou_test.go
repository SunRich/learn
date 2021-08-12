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


func TestMaxArea(t *testing.T) {
	height:=[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := MaxArea(height)
	fmt.Println(result)
	fmt.Println(process2(height))

}

func TestMoveZeroes(t *testing.T) {
	nums:=[]int{0,1,0,3,12}
	MoveZeroes(nums)
	fmt.Println(nums)
}

func TestClimbStairs(t *testing.T) {
	fmt.Println(ClimbStairs(2))
}

func TestThreeSum(t *testing.T) {
	fmt.Println(ThreeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println(ThreeSum([]int{0,0,0,0}))
	fmt.Println(ThreeSum([]int{-2,-5,5,-7,8,4,1,9,-2,-1}))
}

func TestSubsets(t *testing.T)  {
	fmt.Println(subsets([]int{1,2,3}))
}