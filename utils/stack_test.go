package utils

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := new(Stack)
	for i:=0;i<5;i++ {
		stack.Push(i)
	}
	for i:=0;i<5;i++ {
		fmt.Println(stack.Items, stack.Pop())
	}


}
