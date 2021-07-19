package list

import (
	"fmt"
)

type ListNode struct {
	next  *ListNode
	value int
}

func (l *ListNode) Print() {
	head := l
	for {
		fmt.Println(head.value)
		if head.next == nil {
			break
		}
		head = head.next

	}
}

func BuildCycleList() *ListNode {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)
	node1.value = 1
	node2.value = 2
	node3.value = 3
	node4.value = 4
	node5.value = 5

	node1.next = node2
	node2.next = node3
	node3.next = node4

	node4.next = node5

	node5.next = node3
	return node1
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	var temp *ListNode
	for {
		if head == nil {
			break
		}
		temp = head.next
		head.next = pre
		pre = head
		head = temp
	}
	return pre

}

func HasCycleUseMap(head *ListNode) bool {
	if head == nil {
		return false
	}
	headMap := make(map[*ListNode]bool)
	for head != nil {
		if headMap[head] {
			fmt.Println(head.value)
			return true
		} else {
			headMap[head] = true
		}
		head = head.next

	}
	return false
}

//快慢指针拿出对应成环节点
func GetCycleNodeByFS(fast, slow *ListNode) *ListNode {
	if fast == nil || slow == nil {
		return nil
	}
	for fast != nil && slow != nil {
		fast = fast.next
		slow = slow.next
		if fast.value == slow.value {
			return fast
		}

	}
	return nil

}

func HasCycleUseFS(head *ListNode) bool {

	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			fast = head
			for slow != fast {
				slow = slow.next
				fast = fast.next
			}
			if fast == slow {
				fmt.Println(fast.value)
				return true
			}
		}
	}
	return false
}
