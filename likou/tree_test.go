package likou

import (
	"fmt"
	"testing"
)

var head *TreeNode

func init() {
	head = new(TreeNode)
	head.Val = 1
	head.Left = &TreeNode{
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 5,
		},
		Val: 2,
	}
	head.Right = &TreeNode{
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 10,
			},
		},
		Val: 3,
	}
}

func TestLevelOrder(t *testing.T) {
	result := levelOrder(head)
	fmt.Println(result)

	result = zigzagLevelOrder(head)
	fmt.Println(result)
}
