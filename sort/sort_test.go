package sort

import (
	"fmt"
	"testing"
)

func TestList_BubbleSort(t *testing.T) {
	list := NewArraySort([]int64{8,0, 1, 7, 2, 9, 3})
	/*result := list.BubbleSort()
	fmt.Println(result, list.List)

	result = list.SelectSort()
	fmt.Println(result, list.List)
	result = list.InsertSort()
	fmt.Println(result, list.List)*/
	fmt.Println(5>>1)
	result  := list.QuickSort()
	fmt.Println(result, list.List)
}
