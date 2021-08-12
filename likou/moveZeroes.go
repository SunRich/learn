package likou

//283. 移动零
//解法1双指针,右边界遇到非零进行交换
func moveZeroes(nums []int) {
	length := len(nums)
	var left, right int
	for right < length {
		//遇到零不移动
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
		right++
	}
}

//解法2:两个for循环，先把非零的放到前，再把为零的置为空
func moveZeroes2(nums []int) {
	length := len(nums)
	index := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for ; index < length; index++ {
		nums[index]=0
	}
}

func MoveZeroes(nums []int) {
	moveZeroes2(nums)
}
