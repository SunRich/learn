package tree

import "testing"

var head *Tree

func init() {
	head = new(Tree)
	head.value = 1
	head.left = &Tree{
		left: &Tree{
			value: 4,
		},
		right: &Tree{
			value: 5,
		},
		value: 2,
	}
	head.right = &Tree{
		left: &Tree{
			value: 6,
		},
		right: &Tree{
			value: 7,
		},
		value: 3,
	}
}

func TestIteration(t *testing.T) {
	Iteration(head)

}

func TestPreTraverStack(t *testing.T) {
	//Iteration(head)
	PreTraverStack(head)
}
