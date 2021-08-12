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

//
func threeSum3(nums []int) (result [][]int) {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
