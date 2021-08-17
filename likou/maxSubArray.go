package likou

// [-2,1,-3,4,-1,2,1,-5,4]
//分解问题
// 如果当前的i小于i+i-1
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i]+=nums[i-1]
		}
		if nums[i]>max{
			max =nums[i]
		}
	}
	return max
}
