package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	if s.Length != 0 || s.head != nil {
		t.Errorf("NewStack failed. Expected: length=0, head=nil, Got: length=%d, head=%v", s.Length, s.head)
	}
}

func TestPushAndPeek(t *testing.T) {
	s := NewStack[int]()
	s.Push(10)
	s.Push(20)

	peekedItem, ok := s.Peek()
	if !ok || peekedItem != 20 || s.Length != 2 {
		t.Errorf("Peek after pushing failed. Expected: item=20, length=2, Got: item=%d, length=%d, ok=%v", peekedItem, s.Length, ok)
	}
}

func TestPop(t *testing.T) {
	s := NewStack[int]()
	s.Push(10)
	s.Push(20)

	poppedItem, ok := s.Pop()
	if !ok || poppedItem != 20 || s.Length != 1 {
		t.Errorf("Pop failed. Expected: item=20, length=1, Got: item=%d, length=%d, ok=%v", poppedItem, s.Length, ok)
	}

	poppedItem, ok = s.Pop()
	if !ok || poppedItem != 10 || s.Length != 0 {
		t.Errorf("Pop failed. Expected: item=10, length=0, Got: item=%d, length=%d, ok=%v", poppedItem, s.Length, ok)
	}

	// Test popping from an empty stack
	poppedItem, ok = s.Pop()
	if ok || poppedItem != 0 {
		t.Errorf("Pop from empty stack failed. Expected: item=0, ok=false, Got: item=%d, ok=%v", poppedItem, ok)
	}
}

func TestStack(t *testing.T) {
	t.Run("NewStack", TestNewStack)
	t.Run("PushAndPeek", TestPushAndPeek)
	t.Run("Pop", TestPop)
}
