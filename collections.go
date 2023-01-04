package goterator

import "errors"

type SliceCollection[E any] struct {
	elements []*E
}

func (a *SliceCollection[E]) CreateIterator() Iterator[E] {
	return &SliceIterator[E]{index: 0, elements: a.elements}
}

func (a *SliceCollection[E]) Add(element E) {
	a.elements = append(a.elements, &element)
}

func (a *SliceCollection[E]) Remove(index int) error {
	if a.IsEmpty() {
		return errors.New("collection is empty")
	}
	if index > len(a.elements) || index < 0 {
		return errors.New("index out of range")
	}
	a.elements = append(a.elements[0:index], a.elements[index+1:]...)
	return nil
}

func (a *SliceCollection[E]) IsEmpty() bool {
	if len(a.elements) > 0 {
		return false
	}
	return true
}

type SliceIterator[E any] struct {
	index          int
	currentElement *E
	elements       []*E
}

func (a *SliceIterator[E]) HasNext() bool {
	if a.index <= len(a.elements)-1 {
		return true
	}
	return false
}

func (a *SliceIterator[E]) GetNext() *E {
	current := a.elements[a.index]
	a.index += 1
	a.currentElement = current
	return current
}

func (a *SliceIterator[E]) Current() *E {
	return a.currentElement
}
