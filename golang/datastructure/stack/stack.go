package stack

type Stack struct {
	arr []interface{}
}

func (s *Stack) Push(obj interface{}) {
	s.arr = append(s.arr, obj)
}

func (s *Stack) Pop() interface{} {
	length := s.Len()
	if length == 0 {
		panic("Stack is empty")
	}

	obj := s.arr[length-1]
	s.arr = s.arr[:length-1]
	return obj
}

func (s *Stack) Len() int {
	return len(s.arr)
}
