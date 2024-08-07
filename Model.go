package HyperUtility

import "fmt"

var (
	ErrEmpty = fmt.Errorf("empty")
)

// DataStructure define a queue and stack interface containing all function to be implemented by a Queue or Stack
// A Queue will implement the FIFO (First In First Out) logic.
// A Stack will implement the FILO (First In Last Out) logic.
type DataStructure interface {
	// Push will add element into this collection
	Push(interface{})

	// Pop will return an element from this collection, the element will be removed.
	Pop() (interface{}, error)

	// Peek will return an element from this collection, the element will NOT be removed.
	Peek() (interface{}, error)

	// Size counts all elements in this collection.
	Size() int

	// Clear remove ALL element from this collection.
	Clear()
}
