package slidingwindow

type Queue struct {
	Items []int
}

func NewQueue(len int) *Queue {
	return &Queue{Items: make([]int, len, len)}
}

func (q *Queue) Empty() bool {
	if len(q.Items) == 0 {
		return true
	}
	return false
}

func (q *Queue) Push(item int) {
	q.Items = append(q.Items, item)
}

func (q *Queue) PopBack() (item int) {
	if q.Empty() {
		item = -1
		return
	}
	item = q.Items[len(q.Items)-1]
	q.Items = q.Items[:len(q.Items)-1]
	return
}

func (q *Queue) PopFirst() (item int) {
	if q.Empty() {
		item = -1
		return
	}
	item = q.Items[0]
	q.Items = q.Items[1:]
	return
}

func (q *Queue) PeekFirst() (item int) {
	if q.Empty() {
		item = -1
		return
	}
	item = q.Items[0]
	return
}

func (q *Queue) PeekBack() (item int) {
	if q.Empty() {
		item = -1
		return
	}
	item = q.Items[len(q.Items)-1]
	return
}

// [1,3,-1,-3,5,3,6,7], k = 3
func MaxSlidingWindow(nums []int, k int) (result []int) {
	var length = len(nums)
	if length == 0 || k == 0 {
		return []int{}
	}
	var left, right int
	var dequeList = NewQueue(k)
	var i = 0
	for right < length && left <= length-k {
		//先尝试右边从尾部入栈
		for !dequeList.Empty() && nums[dequeList.PeekBack()] < nums[i] {
			dequeList.PopBack()
		}
		dequeList.Push(i)
		right++
		//左边移动，判断当前头部节点是否过期
		index := dequeList.PeekFirst()
		for ; index < left && !dequeList.Empty(); index = dequeList.PeekFirst() {
			dequeList.PopFirst()
		}
		if right-k == left {
			result = append(result, nums[index])
			left++
		}
		i++
	}
	return
}

func MaxSlidingWindow2(nums []int, k int) []int {
	var result []int
	var stack []int
	stack = make([]int, k, k)
	for i, num := range nums {
		//先尝试右边从尾部入栈
		for len(stack) >0 && num >= nums[stack[len(stack)-1]]  {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if i-k+1 > stack[0] {
			stack = stack[1:]
		}
		if i+1>=k{
			result = append(result, nums[stack[0]])
		}
	}
	return result
}
