package containers

type listNode[T any] struct {
	next  *listNode[T]
	prev  *listNode[T]
	value T
}

type List[T any] struct {
	head *listNode[T]
	tail *listNode[T]
	size int64
}

func (l List[T]) Front() T {
	return l.head.value
}

func (l List[T]) Back() T {
	return l.tail.value
}

func (l List[T]) Size() int64 {
	return l.size
}

func (l List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) PushFront(value T) {

	defer func() { l.size++ }()

	if l.head == nil {
		l.head = &listNode[T]{value: value}
		l.tail = l.head
		return
	}

	if l.head == l.tail {
		newNode := &listNode[T]{value: value}
		l.head = newNode
		l.head.next = l.tail
		l.tail.prev = l.head
		return
	}

	newNode := &listNode[T]{value: value}
	temp := l.head
	newNode.next = temp
	temp.prev = newNode
	l.head = newNode
}

func (l *List[T]) PushBack(value T) {

	defer func() { l.size++ }()

	if l.head == nil {
		l.head = &listNode[T]{value: value}
		l.tail = l.head
		return
	}

	if l.head == l.tail {
		newNode := &listNode[T]{value: value}
		l.tail = newNode
		l.head.next = l.tail
		l.tail.prev = l.head
		return
	}

	newNode := &listNode[T]{value: value}
	temp := l.tail
	newNode.prev = temp
	temp.next = newNode
	l.tail = newNode
}

func (l *List[T]) PopFront() {

	defer func() {
		if l.size > 0 {
			l.size--
		}
	}()

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		return
	}

	l.head = l.head.next
}

func (l *List[T]) PopBack() {

	defer func() {
		if l.size > 0 {
			l.size--
		}
	}()

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		return
	}

	l.tail = l.tail.prev
}
