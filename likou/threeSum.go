package likou

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

//解法1 先排序后双指针,主要变量k,left,right
func threeSum(nums []int) (result [][]int) {
	length := len(nums)
	if length < 3 {
		return
	}
	sort.Ints(nums)
	for k := 0; k < length-2; k++ {
		//找到的情况
		n1 := nums[k]
		if n1 > 0 {
			break
		}
		if k > 0 && n1 == nums[k-1] {
			continue
		}
		left, right := k+1, length-1

		for left < right {
			leftNum, rightNum := nums[left], nums[right]

			if n1+leftNum+rightNum == 0 {
				result = append(result, []int{n1, leftNum, rightNum})
				for left < right && nums[left] == leftNum {
					left++

				}
				for left < right && nums[right] == rightNum {
					right--
				}
			} else if n1+leftNum+rightNum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
