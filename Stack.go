package HyperUtility

// NewStack create a new implementation of Stack.
func NewStack() DataStructure {
	return &stackImpl{
		data: make([]interface{}, 0),
	}
}

// stackImpl is a new implementation of Stack.
type stackImpl struct {
	data []interface{}
}

// Push add an item into this stack.
func (impl *stackImpl) Push(i interface{}) {
	impl.data = append(impl.data, i)
}

// Pop take the Stack head to be returned. The length of the
// stack will be decreased by 1.
func (impl *stackImpl) Pop() (interface{}, error) {
	i, err := impl.Peek()
	if err != nil {
		return nil, err
	}
	impl.data = impl.data[:len(impl.data)-1]
	return i, nil
}

// Peek take the Stack head to be returned. The length of the
// stack remain.
func (impl *stackImpl) Peek() (interface{}, error) {
	if len(impl.data) == 0 {
		return nil, ErrEmpty
	}
	i := impl.data[len(impl.data)-1]
	return i, nil
}

// Size get number of elements count in this Stack
func (impl *stackImpl) Size() int {
	return len(impl.data)
}

// Clear will empty this Stack
func (impl *stackImpl) Clear() {
	impl.data = make([]interface{}, 0)
}
