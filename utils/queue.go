package utils

type Queue struct {
	Items []int
}

func (q *Queue) NewQueue() *Queue {
	return &Queue{Items: make([]int, 0)}
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

func (q *Queue) Pop() (item int) {
	if q.Empty() {
		item = -1
		return
	}
	item = q.Items[0]
	q.Items = q.Items[1:]
	return
}

func (q *Queue) Peek() (item int) {
	if q.Empty() {
		item = -1
		return
	}
	item = q.Items[0]
	return
}

func (q *Queue) PeekFirst() (item int) {
	if q.Empty() {
		item = -1
		return
	}
	item = q.Items[len(q.Items)-1]
	return
}
