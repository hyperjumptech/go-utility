package HyperUtility

import "testing"

func TestStackImpl(t *testing.T) {
	// Create new stack.
	s := NewStack()
	// Check if the new stack is empty
	if s.Size() != 0 {
		t.Fail()
	}

	// Add int 1 into stack.
	s.Push(1)
	// Check if the stack has 1 element
	if s.Size() != 1 {
		t.Fail()
	}

	// Add int 2 into stack
	s.Push(2)
	// Check if the stack has 2 elements
	if s.Size() != 2 {
		t.Fail()
	}

	// Peek into the stack and check it content
	if p, err := s.Peek(); err != nil {
		t.Fail()
	} else if p.(int) != 2 {
		t.Fail()
	}
	// Check after peek, the size must not change
	if s.Size() != 2 {
		t.Fail()
	}

	// Pop form the stack and check it content
	if p, err := s.Pop(); err != nil {
		t.Fail()
	} else if p.(int) != 2 {
		t.Fail()
	}
	// Check after pop, the size must be decreased
	if s.Size() != 1 {
		t.Fail()
	}

	// Peek again the value after pop.
	if p, err := s.Peek(); err != nil {
		t.Fail()
	} else if p.(int) != 1 {
		t.Fail()
	}
	// Check after peek on the last content.
	if s.Size() != 1 {
		t.Fail()
	}

	// Pop again the last value
	if p, err := s.Pop(); err != nil {
		t.Fail()
	} else if p.(int) != 1 {
		t.Fail()
	}
	// Check again that the stack his empty
	if s.Size() != 0 {
		t.Fail()
	}

	// Test to peek on the empty stack should fail.
	if _, err := s.Peek(); err == nil {
		t.Fail()
	}

	// Test to pop on the empty stack should fail.
	if _, err := s.Pop(); err == nil {
		t.Fail()
	}

	// Push 3 values.
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Check the size of the stack after adding 3 values.
	if s.Size() != 3 {
		t.Fail()
	}

	// Clear the stack
	s.Clear()

	// Make sure that the stack is emptied after clear
	if s.Size() != 0 {
		t.Fail()
	}
}
