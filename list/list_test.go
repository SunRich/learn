package list

import (
	"fmt"
	"testing"
)

var head *ListNode

func init() {
	head = new(ListNode)
	head.value = 1
	head.next = &ListNode{
		value: 2,
		next: &ListNode{
			value: 3,
			next: &ListNode{
				value: 4,
				next: &ListNode{
					value: 5,
				},
			},
		},
	}
}

func TestReverseList(t *testing.T) {

	head.Print()
	newHead := ReverseList(head)
	fmt.Println("\n","after")

	newHead.Print()

}

func TestHasCycleUseMap(t *testing.T)  {
	fmt.Println(HasCycleUseMap(head))
	//fmt.Println(HasCycleUseMap(BuildCycleList()))
	fmt.Println(HasCycleUseFS(BuildCycleList()))

}
