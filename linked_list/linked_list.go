package linkedlist

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

func (l *LinkedList) Head() *Node {
	return l.head
}

// Insert 最後尾にNodeを追加する
func (l *LinkedList) Append(data int) {
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

// Insert 先頭にNodeを追加する
func (l *LinkedList) Insert(data int) {
	newNode := NewNode(data)
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) Remove(data int) {
	currentNode := l.head
	if currentNode != nil && currentNode.data == data {
		l.head = currentNode.next
		currentNode = nil
		return
	}

	var previousNode *Node
	for currentNode != nil && currentNode.data != data {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	if currentNode == nil {
		return
	}

	previousNode.next = currentNode.next
}

func (l *LinkedList) Print(writer io.Writer) {
	iterator := l.head
	for iterator != nil {
		fmt.Fprintf(writer, "%v ", iterator.data)
		iterator = iterator.next
	}
}

func (l *LinkedList) ReverseIterative() {
	var prevNode *Node
	currentNode := l.head
	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	l.head = prevNode
}

func (l *LinkedList) ReverseRecursive() {
	l.head = reverseRecursiveHelper(nil, l.head)
}

func reverseRecursiveHelper(previous, current *Node) *Node {
	if current == nil {
		return previous
	}

	nextNode := current.next
	current.next = previous
	previous = current
	current = nextNode

	return reverseRecursiveHelper(previous, current)
}

func (l *LinkedList) ReverseEven() {
	var prevNode *Node
	currentNode := l.head
	for !reflect.ValueOf(currentNode).IsNil() {
		if currentNode.data%2 == 0 {
			beforeNode := prevNode
			startNode := currentNode
			endNode, nextNode := reverseEvenHelper(currentNode)
			beforeNode.next = endNode
			startNode.next = nextNode
			currentNode = nextNode
			prevNode = startNode
		} else {
			prevNode = currentNode
			currentNode = currentNode.next
		}
	}
}

func reverseEvenHelper(startNode *Node) (endNode *Node, nextNode *Node) {
	prevNode := startNode
	currentNode := startNode.next
	for currentNode != nil && currentNode.data%2 == 0 {
		nextNode := currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	return prevNode, currentNode
}

func Run() {
	l := NewLinkedList()
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	fmt.Println(strings.Repeat("#", 10))
	l.Print(os.Stdout)
	l.ReverseIterative()
	l.Print(os.Stdout)
	l.ReverseRecursive()
	l.Print(os.Stdout)
}
