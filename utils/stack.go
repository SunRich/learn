package utils

type Stack struct {
	Items []interface{}
}

func NewStack() (s *Stack) {
	return new(Stack)
}

func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() (item interface{}) {
	if s.IsEmpty() {
		return
	}
	item = s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}
