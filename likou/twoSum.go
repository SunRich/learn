package likou

func twoSum(nums []int, target int) []int {
	valueMap := make(map[int]int)

	for index, value := range nums {
		if v, ok := valueMap[target-value]; ok {

			return []int{v,index}
		} else {
			valueMap[value] = index
		}
	}
	return []int{}
}

