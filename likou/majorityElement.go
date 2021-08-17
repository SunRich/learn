package likou

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
