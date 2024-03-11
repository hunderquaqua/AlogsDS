package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	// Create a new queue
	q := NewQueue[int]()

	// Enqueue items
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Dequeue all items and check if they are dequeued in the correct order
	expected := []int{10, 20, 30}
	for _, exp := range expected {
		item, ok := q.Deque()
		if !ok || item != exp {
			t.Errorf("Dequeued item is incorrect. Expected: %d, Got: %d", exp, item)
		}
	}

	// Try to dequeue from an empty queue
	item, ok := q.Deque()
	if ok || item != 0 {
		t.Errorf("Dequeued from empty queue should return 0, false. Got: %d, %v", item, ok)
	}

	// Enqueue two elements
	q.Enqueue(100)
	q.Enqueue(200)

	// Peek at the front item
	frontItem, ok := q.Peek()
	if !ok || frontItem != 100 {
		t.Errorf("Peek failed. Expected: 100, Got: %d", frontItem)
	}
}

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q.Length != 0 || q.head != nil || q.tail != nil {
		t.Errorf("NewQueue failed. Expected: length=0, head=nil, tail=nil, Got: length=%d, head=%v, tail=%v", q.Length, q.head, q.tail)
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	if q.Length != 2 {
		t.Errorf("Enqueue failed. Expected length=2, Got: length=%d", q.Length)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)

	item, ok := q.Deque()
	if !ok || item != 10 || q.Length != 1 {
		t.Errorf("Dequeue failed. Expected: item=10, length=1, Got: item=%d, length=%d, ok=%v", item, q.Length, ok)
	}

	item, ok = q.Deque()
	if !ok || item != 20 || q.Length != 0 {
		t.Errorf("Dequeue failed. Expected: item=20, length=0, Got: item=%d, length=%d, ok=%v", item, q.Length, ok)
	}

	item, ok = q.Deque()
	if ok || item != 0 {
		t.Errorf("Dequeue from empty queue failed. Expected: item=0, ok=false, Got: item=%d, ok=%v", item, ok)
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)

	item, ok := q.Peek()
	if !ok || item != 10 || q.Length != 2 {
		t.Errorf("Peek failed. Expected: item=10, length=2, Got: item=%d, length=%d, ok=%v", item, q.Length, ok)
	}
}
