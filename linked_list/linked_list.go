package linkedlist

import (
	"fmt"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{
		data: data,
		next: nil,
	}
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
	}
}

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

// Insert　最後尾にNodeを追加する
func (l *LinkedList[T]) Append(data T) {
	newNode := NewNode(data)
	if l.head == nil {
		l.head = newNode
		return
	}

	iterator := l.head
	for iterator.next != nil {
		iterator = iterator.next
	}
	iterator.next = newNode
}

// Insert　先頭にNodeを追加する
func (l *LinkedList[T]) Insert(data T) {
	newNode := NewNode(data)
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList[T]) Remove(data T) {
	currentNode := l.head
	if currentNode != nil && currentNode.data == data {
		l.head = currentNode.next
		currentNode = nil
		return
	}

	var previousNode *Node[T]
	for currentNode != nil && currentNode.data != data {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	if currentNode == nil {
		return
	}

	previousNode.next = currentNode.next
}

func (l *LinkedList[T]) Print() {
	iterator := l.head
	for iterator != nil {
		fmt.Printf("%v ", iterator.data)
		iterator = iterator.next
	}
	fmt.Println()
}

func Run() {
	l := NewLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Insert(0)
	l.Print()
	l.Remove(3)
	l.Print()
	l.Remove(0)
	l.Print()
	l.Remove(4)
	l.Print()
	l.Remove(100)
	l.Print()
}
