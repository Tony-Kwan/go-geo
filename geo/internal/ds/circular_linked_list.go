package ds

import (
	"fmt"
	"strings"
)

type CircularLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

type Node struct {
	Elem interface{}
	Prev *Node
	Next *Node
}

type Iterator struct {
	list  *CircularLinkedList
	index int
	node  *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (l *CircularLinkedList) Add(e interface{}) *CircularLinkedList {
	node := &Node{Elem: e}
	if l.Head == nil && l.Tail == nil {
		node.Prev = node
		node.Next = node
		l.Head = node
		l.Tail = node
	} else {
		node.Prev = l.Tail
		node.Next = l.Tail.Next
		l.Tail.Next = node
		l.Head.Prev = node
		l.Tail = node
	}
	l.size++
	return l
}

func (l *CircularLinkedList) Insert(node *Node, e interface{}) *CircularLinkedList {
	if node == nil {
		return l
	}
	newNode := &Node{
		Elem: e,
		Prev: node,
		Next: node.Next,
	}
	if node.Next != nil {
		node.Next.Prev = newNode
	}
	node.Next = newNode
	l.size++
	return l
}

func (l *CircularLinkedList) Remove(idx int) *Node {
	if idx >= l.size || idx < 0 {
		return nil
	}
	if l.size == 1 {
		l.size = 0
		l.Head = nil
		l.Tail = nil
		return nil
	}
	node := l.Head
	for i := 0; i != idx; i, node = i+1, node.Next {
	}
	return l.RemoveNode(node)
}

func (l *CircularLinkedList) RemoveNode(node *Node) *Node {
	if node.Prev == nil || node.Next == nil {
		return nil
	}
	if l.size == 0 {
		return nil
	}
	if l.size == 1 {
		if l.Head == node {
			l.Head = nil
			l.Tail = nil
			node.Prev = nil
			node.Next = nil
			l.size--
		}
		return nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		if l.Head == node {
			l.Head = node.Next
		}
		ret := node.Next
		node.Prev = nil
		node.Next = nil
		l.size--
		return ret
	}
}

func (l *CircularLinkedList) Size() int {
	return l.size
}

func (l *CircularLinkedList) Iterator() *Iterator {
	return &Iterator{list: l, index: -1, node: nil}
}

func (l *CircularLinkedList) String() string {
	s := "CircularLinkedList["
	ss := make([]string, 0, l.size)
	for it := l.Iterator(); it.Next(); {
		ss = append(ss, fmt.Sprintf("%v", it.Value()))
	}
	return s + strings.Join(ss, ", ") + "]"
}

func (it *Iterator) Next() bool {
	if it.index < it.list.size {
		it.index++
	}
	if it.index >= it.list.size {
		return false
	}
	if it.index == 0 {
		it.node = it.list.Head
	} else {
		it.node = it.node.Next
	}
	return true
}

func (it *Iterator) Value() interface{} {
	return it.node.Elem
}

func (it *Iterator) Index() int {
	return it.index
}

func (it *Iterator) Node() *Node {
	return it.node
}
