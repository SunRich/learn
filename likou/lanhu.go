package likou

func DeleteEvenNumber(nums []int) (result []int) {
	i,j:=0,0
	for i<len(nums)-1 && j<len(nums) {
		if nums[i]%2==0  {
			if nums[j]%2!=0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			}else{
				j++
			}

		}else{
			i++
			j++
		}

	}
	result = nums[:i]
	return
}