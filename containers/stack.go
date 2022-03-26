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
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Pop() {
	s.data = s.data[:len(s.data)-1]
}
