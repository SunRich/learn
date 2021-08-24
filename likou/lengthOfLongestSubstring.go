package likou

//3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	doubleMap := make([]int, 128)
	var start, length, nums int
	for i, v := range []byte(s) {
		if doubleMap[v] > start {
			start = doubleMap[v]
		}
		nums = i - start + 1
		if nums > length {
			length = nums
		}
		doubleMap[v] = i + 1
	}
	return length
}
