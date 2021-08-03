package monotonic_stack

import (
	"fmt"
	"testing"
)

func TestSortedStack(t *testing.T) {
	result := FindNumbersLeftRightFirstMax([]int{1, 5,3,6,4})
	fmt.Println(FindLeftFirstMaxByLeftRight(result))
	fmt.Println(result)
}
