package ds

import (
	"fmt"
	"testing"
)

var l *CircularLinkedList

func TestMain(m *testing.M) {
	l = NewCircularLinkedList()
	l.Add(1).Add(2).Add(3)
	m.Run()
}

func TestCircularLinkedList_Add(t *testing.T) {
	node := l.Head
	for i := 0; i < l.size; i++ {
		t.Log(node.Elem)
		node = node.Next
	}
}

func TestCircularLinkedList_RemoveNode(t *testing.T) {
	node := l.Head.Next.Next.Next.Next.Next.Next.Next
	fmt.Println(l.String())
	node = l.RemoveNode(node)
	fmt.Println(l.String())
	node = l.RemoveNode(node)
	fmt.Println(l.String())
}

func TestCircularLinkedList_Remove(t *testing.T) {
	l.Remove(2)
	fmt.Println(l.String())
}
