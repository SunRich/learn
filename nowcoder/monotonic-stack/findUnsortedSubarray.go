package monotonic_stack

import "fmt"

type LeftRight struct {
	Left  int
	Right int
	Index int
}
// FindUnsortedSubarray nums = [2,6,4,8,10,9,15]
func FindUnsortedSubarray(nums []int) int {
	left:=0
	right:=-1
	leftMax:=nums[0]
	rightMix:=nums[len(nums)-1]
	for i:=0;i<len(nums);i++{
		if leftMax<=nums[i] {
			leftMax =nums[i]
		}else{
			right=i
		}
		if rightMix>=nums[len(nums)-1-i] {
			rightMix =nums[len(nums)-1-i]
		}else{
			left=len(nums)-1-i
		}
	}
	fmt.Println(right,left)
	return right-left+1
}

func FindLeftFirstMaxByLeftRight(leftRight []LeftRight) (result []int) {
	result = make([]int, len(leftRight))
	for _, v := range leftRight {
		result[v.Index] = v.Left
	}
	return
}

// FindNumbersLeftRightFirstMax 通过单调栈获取数组元素左右两边大于该元素的值
func FindNumbersLeftRightFirstMax(nums []int) []LeftRight {
	result := make([]LeftRight, 0)
	if len(nums) == 0 {
		return result
	}
	stack := make([]int, 0)
	for i, num := range nums {
		//弹出的逻辑
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			leftRight := LeftRight{
				Right: i,
				Index: stack[len(stack)-1],
			}
			if len(stack) >= 2 {
				leftRight.Left = stack[len(stack)-2]
			} else {
				leftRight.Left = -1
			}
			result = append(result, leftRight)
			stack = stack[:len(stack)-1]

		}
		stack = append(stack, i)
	}
	//剩余的弹出
	for len(stack) > 0 {
		leftRight := LeftRight{
			Right: -1,
			Index: stack[len(stack)-1],
		}
		if len(stack) >= 2 {
			leftRight.Left = stack[len(stack)-2]
		} else {
			leftRight.Left = -1
		}
		result = append(result, leftRight)
		stack = stack[:len(stack)-1]
	}
	return result
}
