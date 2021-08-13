package kmp

func FindIndexOf(str, match string) (index int) {
	//异常判断

	var x = 0
	var y = 0
	strByte := []byte(str)
	matchByte := []byte(match)

	next := GetNextArray(matchByte)
	for x < len(strByte) && y < len(matchByte) {
		if strByte[x] == matchByte[y] {
			x++
			y++
		} else if next[y] == -1 {
			x++
		} else {
			y = next[y]
		}
	}
	if y == len(matchByte) {
		return x - y
	}
	return -1
}

func GetNextArray(str []byte) (next []int) {
	next = make([]int, len(str))
	if len(str) == 0 {
		return
	}
	if len(str) == 1 {
		next = []int{-1}
		return
	}
	next[0] = -1
	next[1] = 0
	var p = 2
	var cn = 0
	for p < len(str) {
		if str[p-1] == str[cn] {
			next[p] = cn + 1
			p = p + 1
			cn = cn + 1
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[p] = 0
			p++
		}
	}
	return
}
