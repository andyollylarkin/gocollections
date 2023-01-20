package gocollections

import (
	"errors"
	"sync"
)

type SliceCollection[E any] struct {
	elements []*E
	m        sync.Mutex
}

func (a *SliceCollection[E]) CreateIterator() Iterator[E] {
	return &SliceIterator[E]{index: 0, elements: a.elements}
}

func (a *SliceCollection[E]) Add(element E) {
	a.m.Lock()
	defer a.m.Unlock()
	a.elements = append(a.elements, &element)
}

func (a *SliceCollection[E]) Remove(index int) error {
	a.m.Lock()
	defer a.m.Unlock()
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
	m              sync.Mutex
}

func (a *SliceIterator[E]) HasNext() bool {
	a.m.Lock()
	defer a.m.Unlock()
	if a.index <= len(a.elements)-1 {
		return true
	}
	return false
}

func (a *SliceIterator[E]) GetNext() *E {
	a.m.Lock()
	defer a.m.Unlock()
	if a.index+1 > len(a.elements)-1 {
		panic("index out of range. Index: " + string(rune(a.index+1)))
	}
	current := a.elements[a.index]
	a.index += 1
	a.currentElement = current
	return current
}

func (a *SliceIterator[E]) Current() *E {
	return a.currentElement
}

type ListNode[E any] struct {
	Val      E
	next     *ListNode[E]
	previous *ListNode[E]
	head     *ListNode[E]
	tail     *ListNode[E]
	len      int
	m        sync.Mutex
}

func NewListNode[E any](val E) *ListNode[E] {
	node := new(ListNode[E])
	node.Val = val
	node.head = node
	node.tail = node
	node.len = 1
	return node
}

func (l *ListNode[E]) CreateIterator() Iterator[ListNode[E]] {
	return &LinkedListIterator[E]{currentElement: l.head}
}

func (l *ListNode[E]) Add(element E) {
	l.m.Lock()
	defer l.m.Unlock()
	newNode := NewListNode(element)
	newNode.next = nil
	newNode.previous = l.tail
	l.tail.next = newNode
	l.tail = newNode
	l.len += 1
}

func (l *ListNode[E]) IsEmpty() bool {
	return l.len == 0
}

func (l *ListNode[E]) Remove(index int) error {
	l.m.Lock()
	defer l.m.Unlock()
	if index > l.len {
		return errors.New("no an element with that index")
	}
	current := l.head
	for i := 0; i <= index; i++ {
		current = current.next
	}
	current.previous.next = current.next
	return nil
}

type LinkedListIterator[E any] struct {
	currentElement *ListNode[E]
	m              sync.Mutex
}

func (l *LinkedListIterator[E]) HasNext() bool {
	l.m.Lock()
	defer l.m.Unlock()
	return l.currentElement.next != nil
}

func (l *LinkedListIterator[E]) GetNext() *ListNode[E] {
	l.m.Lock()
	defer l.m.Unlock()
	if l.currentElement.next != nil {
		l.currentElement = l.currentElement.next
		return l.currentElement
	}
	return l.currentElement
}

func (l *LinkedListIterator[E]) Current() *ListNode[E] {
	return l.currentElement
}
