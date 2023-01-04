package goterator

// Iterator base interface for iterators
type Iterator[T any] interface {
	HasNext() bool
	GetNext() *T
	Current() *T
}

// Collection base interface for collections
type Collection[T any] interface {
	CreateIterator() Iterator[T]
	Add(element T)
	IsEmpty() bool
	Remove(index int) error
}

// ForEach do iteration over struct which implement Iterator interface
func ForEach[T any](iterator Iterator[T], forBody func(element *T)) {
	for iterator.HasNext() {
		forBody(iterator.GetNext())
	}
}
