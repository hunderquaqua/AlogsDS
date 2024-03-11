package stack

type Node[T any] struct {
	value T
	prev  *Node[T]
}

type Stack[T any] struct {
	Length int
	head   *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Length: 0,
		head:   nil,
	}
}

func (s *Stack[T]) Push(item T) {
	newNode := &Node[T]{value: item, prev: s.head}
	s.head = newNode
	s.Length++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Length == 0 {
		var nilVal T
		return nilVal, false
	}

	head := s.head
	s.head = s.head.prev
	s.Length--

	return head.value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.head == nil {
		var nilVal T
		return nilVal, false
	}

	return s.head.value, true
}
