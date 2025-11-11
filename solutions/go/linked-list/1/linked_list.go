package linkedlist

import "fmt"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
}

func NewList(elements ...interface{}) *List {
	var root = List{}
	for _, el := range elements {
		root.Push(el)
	}
	return &root
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
}

func (l *List) Push(v interface{}) {
	if l.head == nil {
		l.head = &Node{Value: v}
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	newNode := &Node{Value: v, prev: current}
	current.next = newNode
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, fmt.Errorf("list is empty")
	}
	removedValue := l.head.Value
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	return removedValue, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.head == nil {
		return nil, fmt.Errorf("list is empty")
	}
	current := l.head
	if current.next == nil {
		removedValue := current.Value
		l.head = nil
		return removedValue, nil
	}
	for current.next != nil {
		current = current.next
	}
	removedValue := current.Value
	if current.prev != nil {
		current.prev.next = nil
	}
	return removedValue, nil
}

func (l *List) Reverse() {
	var prev *Node
	current := l.head
	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	l.head = prev
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	if l.head == nil {
		return nil
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current
}
