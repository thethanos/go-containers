package containers

type Set[T comparable] struct {
	data map[T]bool
}

func NewSet[T comparable]() Set[T] {

	internal := make(map[T]bool)

	return Set[T]{data: internal}
}

func (s *Set[T]) Insert(value T) {

	if s.data == nil {
		panic("set is uninitialized")
	}

	s.data[value] = true
}

func (s *Set[T]) Remove(value T) {

	if s.data == nil {
		panic("set is uninitialized")
	}

	delete(s.data, value)
}

func (s Set[T]) Contains(value T) bool {

	if s.data == nil {
		panic("set is uninitialized")
	}

	return s.data[value]
}

func (s Set[T]) Size() int {
	return len(s.data)
}

func (s Set[T]) Empty() bool {
	return len(s.data) == 0
}
