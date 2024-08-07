package HyperUtility

// NewQueue create a new implementation of Queue.
func NewQueue() DataStructure {
	return &queueImpl{
		data: make([]interface{}, 0),
	}
}

// queueImpl is a new implementation of Queue.
type queueImpl struct {
	data []interface{}
}

// Push add an item into this queue.
func (impl *queueImpl) Push(i interface{}) {
	impl.data = append(impl.data, i)
}

// Pop take the Queue tail to be returned. The length of the
// queue will be decreased by 1.
func (impl *queueImpl) Pop() (interface{}, error) {
	i, err := impl.Peek()
	if err != nil {
		return nil, err
	}
	impl.data = impl.data[1:]
	return i, nil
}

// Peek take the Queue tail to be returned. The length of the
// queue remain.
func (impl *queueImpl) Peek() (interface{}, error) {
	if impl.Size() == 0 {
		return nil, ErrEmpty
	}
	i := impl.data[0]
	return i, nil
}

// Size get number of elements count in this Queue.
func (impl *queueImpl) Size() int {
	return len(impl.data)
}

// Clear will empty this Queue
func (impl *queueImpl) Clear() {
	impl.data = make([]interface{}, 0)
}
