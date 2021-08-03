package monotonic_stack

type LeftRight struct {
	Left  int
	Right int
	Index int
}

func findUnsortedSubarray(nums []int) int {
	return 0
}

func FindLeftFirstMaxByLeftRight(leftRight []LeftRight) (result []int) {
	result = make([]int, len(leftRight))
	for _, v := range leftRight {
		result[v.Index] = v.Left
	}
	return
}

//通过单调栈获取数组元素左右两边大于该元素的值
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
