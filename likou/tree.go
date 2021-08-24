package likou

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//102. 二叉树的层序遍历
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

//103. 二叉树的锯齿形层序遍历
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

//226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

//98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelp(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelp(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	//注意这里的等于
	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValidBSTHelp(root.Left, min, root.Val) && isValidBSTHelp(root.Right, root.Val, max)
}

//104. 二叉树的最大深度
//跟左树，右树要数据，进行聚合
func maxDepth(root *TreeNode) int {
	return maxDepthHelp(root)
}

func maxDepthHelp(root *TreeNode) int {
	deep := 1
	if root == nil {
		return 0
	}
	leftDeep := maxDepthHelp(root.Left)
	rightDeep := maxDepthHelp(root.Right)
	if leftDeep > rightDeep {
		return deep + leftDeep
	} else {
		return deep + rightDeep
	}
}

//111. 二叉树的最小深度
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	minD := math.MaxInt32
	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}

	return minD + 1
}

//236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//边缘条件，终止条件
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	//递推
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//返回值
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
