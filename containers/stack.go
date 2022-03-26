package containers

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s Stack[T]) Top() T {

	if s.Empty() {
		panic("cannot retrieve top value, stack is empty")
	}

	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Pop() {

	if s.Empty() {
		panic("cannot pop top value, stack is empty")
	}

	s.data = s.data[:len(s.data)-1]
}
