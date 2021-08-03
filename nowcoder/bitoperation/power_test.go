package bitoperation

import (
	"fmt"
	"testing"
)

func TestTwoPower(t *testing.T) {
	fmt.Println(TwoPower(4))
	fmt.Println(TwoPower(5))

}

func TestFourPower(t *testing.T) {
	fmt.Println((8>>31)&1)
	fmt.Println(FourPower(8))
	fmt.Println(FourPower(9),0X55555555)

}

func TestGetMax(t *testing.T) {
	fmt.Println(GetMax(1,2))
	fmt.Println(GetMax(-1,2))
	fmt.Println(GetMax(-1,-2))
	fmt.Println(GetMax(3,-2))
}
