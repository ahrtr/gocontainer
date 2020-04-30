package collection

// Interface is a type of collection, all containers should implement this interface.
type Interface interface {
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear removes all of the elements from this container.
	Clear()

	// Len is the number of elements in the collection.
	// Len() is also included in sort.Interface. Only golang 1.14 supports embedding of Interfaces with overlapping method sets,
	// so let add it in this interface in the future.
	//Len() int
}
