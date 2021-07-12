package sort

import (
	"fmt"
	"testing"
)

func TestList_BubbleSort(t *testing.T) {
	list := NewArraySort([]int64{0, 1, 7, 2, 9, 3, 8})
	/*result := list.BubbleSort()
	fmt.Println(result, list.List)

	result = list.SelectSort()
	fmt.Println(result, list.List)
	result = list.InsertSort()
	fmt.Println(result, list.List)*/

	result  := list.QuickSort()
	fmt.Println(result, list.List)
}
