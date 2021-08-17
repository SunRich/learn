package likou

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		level := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			head := queue[0]
			queue = queue[1:]
			level = append(level, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		result = append(result, level)

	}
	return result
}

var res [][]int

func process1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res = make([][]int, 0)
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	levelNum := 1
	for len(queue) > 0 {
		level := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			head := queue[0]
			queue = queue[1:]
			level = append(level, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		if levelNum&1 == 0 {
			left, right := 0, len(level)-1
			for left < right {
				level[left], level[right] = level[right], level[left]
				left++
				right--
			}
		}
		result = append(result, level)
		levelNum++

	}
	return result
}
