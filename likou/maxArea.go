package likou

//11. 盛最多水的容器
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49

//解法双指针
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var max int
	for i < j {
		cur := (j - i) * mix(height[i], height[j])
		if cur > max {
			max = cur
		}
		//谁小谁移动
		if height[i] < height[j] {
			i++
		} else {
			j-- //注意这里是--
		}
	}
	return max
}

//暴力破解
func process2(height []int) int {
	length := len(height)
	var max int
	for i := 0; i < length-1; i++ {
		for j := i + 1; j< length; j++ {
			cur := (j - i) * mix(height[i], height[j])
			if cur > max {
				max = cur
			}
		}
	}
	return max
}


func MaxArea(height []int) int {
	return maxArea(height)
}
