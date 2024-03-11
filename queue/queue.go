package queue

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue[T]) Enqueue(item T) {
	newNode := &Node[T]{value: item}

	if q.head == nil {
		q.head = newNode
	} else {
		q.tail.next = newNode
	}
	q.tail = newNode
	q.Length++
}

func (q *Queue[T]) Deque() (T, bool) {
	if q.head == nil {
		var nilVal T
		return nilVal, false
	}

	q.Length--
	head := q.head
	q.head = q.head.next
	head.next = nil

	return head.value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		var nilVal T
		return nilVal, false
	}

	return q.head.value, true
}
