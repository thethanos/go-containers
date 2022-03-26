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

	if l.Empty() {
		panic("cannot retrieve front value, list is empty")
	}

	return l.head.value
}

func (l List[T]) Back() T {

	if l.Empty() {
		panic("cannot retrieve back value, list is empty")
	}

	return l.tail.value
}

func (l List[T]) Size() int64 {
	return l.size
}

func (l List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) PushFront(value T) {

	if l.head == nil {
		l.head = &listNode[T]{value: value}
		l.tail = l.head
		l.size++
		return
	}

	if l.head == l.tail {
		newNode := &listNode[T]{value: value}
		l.head = newNode
		l.head.next = l.tail
		l.tail.prev = l.head
		l.size++
		return
	}

	newNode := &listNode[T]{value: value}
	temp := l.head
	newNode.next = temp
	temp.prev = newNode
	l.head = newNode
	l.size++
}

func (l *List[T]) PushBack(value T) {

	if l.head == nil {
		l.head = &listNode[T]{value: value}
		l.tail = l.head
		l.size++
		return
	}

	if l.head == l.tail {
		newNode := &listNode[T]{value: value}
		l.tail = newNode
		l.head.next = l.tail
		l.tail.prev = l.head
		l.size++
		return
	}

	newNode := &listNode[T]{value: value}
	temp := l.tail
	newNode.prev = temp
	temp.next = newNode
	l.tail = newNode
	l.size++
}

func (l *List[T]) PopFront() {

	if l.Empty() {
		panic("cannot pop front value, list is empty")
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return
	}

	l.head = l.head.next
	l.size--
}

func (l *List[T]) PopBack() {

	if l.Empty() {
		panic("cannot pop back value, list is empty")
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return
	}

	l.tail = l.tail.prev
	l.size--
}
