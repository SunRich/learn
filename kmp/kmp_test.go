package kmp

import (
	"fmt"
	"testing"
)

func TestGetNextArray(t *testing.T) {
	fmt.Println(GetNextArray([]byte("abacabacabc")))
	fmt.Println(GetNextArray([]byte("abcabc")))

}

func TestFindIndexOf(t *testing.T) {
	fmt.Println(FindIndexOf("abacabadcabc","dca"))
}