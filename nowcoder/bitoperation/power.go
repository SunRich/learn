package bitoperation


//判断一个数是否为2的幂次方
func TwoPower(number int)bool  {
	if number&(number-1)==0{
		return  true
	}
	return  false
}

func FourPower(number int) bool {
	return TwoPower(number) && number&0X55555555==0
}
