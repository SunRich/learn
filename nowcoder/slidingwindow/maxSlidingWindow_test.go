package slidingwindow

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	/*d:=deque.New(3,3)
	d.PushBack(2)
	fmt.Println(d.PopFront())
	fmt.Println(d.Len())*/

}

func TestMaxSlidingWindow2(t *testing.T) {
	result := MaxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(result)
}
