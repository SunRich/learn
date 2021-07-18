package tree

import (
	"fmt"
	"learn/utils"
)

type Tree struct {
	left  *Tree
	right *Tree
	value int
}

func Iteration(t *Tree) {
	if t == nil {
		return
	}
	fmt.Println(t.value)
	Iteration(t.left)
	Iteration(t.right)


}

func PreTraverStack(t *Tree) {
	stack := utils.NewStack()
	stack.Push(t)
	for !stack.IsEmpty() {
		t, ok := stack.Pop().(*Tree)
		if ok {
			fmt.Println(t.value)
			if t.right != nil {
				stack.Push(t.right)
			}
			if t.left != nil {
				stack.Push(t.left)
			}


		}

	}
}
