package HyperUtility

import "testing"

func TestQueueImpl(t *testing.T) {
	// Create new queue.
	s := NewQueue()
	// Check if the new queue is empty
	if s.Size() != 0 {
		t.Fail()
	}

	// Add int 1 into queue.
	s.Push(1)
	// Check if the queue has 1 element
	if s.Size() != 1 {
		t.Fail()
	}

	// Add int 2 into queue
	s.Push(2)
	// Check if the queue has 2 elements
	if s.Size() != 2 {
		t.Fail()
	}

	// Peek into the queue and check it content
	if p, err := s.Peek(); err != nil {
		t.Fail()
	} else if p.(int) != 1 {
		t.Fail()
	}
	// Check after peek, the size must not change
	if s.Size() != 2 {
		t.Fail()
	}

	// Pop form the queue and check it content
	if p, err := s.Pop(); err != nil {
		t.Fail()
	} else if p.(int) != 1 {
		t.Fail()
	}
	// Check after pop, the size must be decreased
	if s.Size() != 1 {
		t.Fail()
	}

	// Peek again the value after pop.
	if p, err := s.Peek(); err != nil {
		t.Fail()
	} else if p.(int) != 2 {
		t.Fail()
	}
	// Check after peek on the last content.
	if s.Size() != 1 {
		t.Fail()
	}

	// Pop again the last value
	if p, err := s.Pop(); err != nil {
		t.Fail()
	} else if p.(int) != 2 {
		t.Fail()
	}
	// Check again that the queue his empty
	if s.Size() != 0 {
		t.Fail()
	}

	// Test to peek on the empty queue should fail.
	if _, err := s.Peek(); err == nil {
		t.Fail()
	}

	// Test to pop on the empty queue should fail.
	if _, err := s.Pop(); err == nil {
		t.Fail()
	}

	// Push 3 values.
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Check the size of the queue after adding 3 values.
	if s.Size() != 3 {
		t.Fail()
	}

	// Clear the stack
	s.Clear()

	// Make sure that the queue is emptied after clear
	if s.Size() != 0 {
		t.Fail()
	}
}
