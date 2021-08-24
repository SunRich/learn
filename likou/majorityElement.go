package likou

//169. 多数元素
func majorityElement(nums []int) int {
	count:=map[int]int{}
	for _,v:=range nums{
		count[v]++
		if count[v]>len(nums)/2{
			return v
		}
	}
	return 0
}
