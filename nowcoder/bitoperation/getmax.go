package bitoperation

func GetMax1() {

}

func flip(sign int) int {
	return sign ^ 1
}

//获取符号位
func GetSign(number int) int {
	return flip((number >> 31) & 1)
}

func GetMax(numberA, numberB int) int {
	//需要考虑相减后溢出问题
	numberC := numberA - numberB
	signA := GetSign(numberA)
	signB := GetSign(numberB)
	signC := GetSign(numberC)
	diffSign := signA ^ signB //符号相同为1，符号相反为0
	sameSign := flip(diffSign)
	returnA := sameSign*signC + diffSign*signA //符号相同的情况下，signC为正数，符号相反的情况下，SignA为正数
	returnB := flip(returnA)
	return returnA*numberA + returnB*numberB
}
