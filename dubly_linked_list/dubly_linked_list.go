package dllist

import (
	"fmt"
	"io"
	"reflect"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
		prev: nil,
		next: nil,
	}
}

type List struct {
	head *Node
}

func NewList() *List {
	return &List{
		head: nil,
	}
}

// Append 最後尾にNodeを追加
func (l *List) Append(data int) {
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
	newNode.prev = iterator
}

// Insert 先頭にNodeを追加
func (l *List) Insert(data int) {
	newNode := NewNode(data)
	newNode.next = l.head
	l.head = newNode
}

// Remove 引数のdataを保持するNodeを削除。見つからなかったら何もしない。
func (l *List) Remove(data int) {
	if l.head == nil {
		return
	}

	iterator := l.head
	for iterator != nil {
		if iterator.data == data {
			break
		}
		iterator = iterator.next
	}

	if iterator == nil {
		return
	}

	if iterator == l.head {
		l.head = l.head.next
	} else if iterator.next == nil {
		iterator.prev.next = nil
	} else {
		prev := iterator.prev
		next := iterator.next
		prev.next = next
		next.prev = prev
	}
}

func (l *List) ReverseIterative() {
	if l.head == nil {
		return
	}

	var prevNode *Node
	currentNode := l.head
	for currentNode != nil {
		prevNode = currentNode.prev
		currentNode.prev = currentNode.next
		currentNode.next = prevNode
		currentNode = currentNode.prev
	}

	if reflect.ValueOf(prevNode).IsNil() {
		return
	}

	l.head = prevNode.prev
}

func (l *List) ReverseRecursive() {
	l.head = reverseRecursiveHelper(l.head)
}

func reverseRecursiveHelper(currentNode *Node) *Node {
	if currentNode == nil {
		return nil
	}

	prevNode := currentNode.prev
	currentNode.prev = currentNode.next
	currentNode.next = prevNode

	if currentNode.prev != nil {
		return reverseRecursiveHelper(currentNode.prev)
	}

	return currentNode
}

func (l *List) Print(writer io.Writer) {
	iterator := l.head

	isFirst := true
	for !reflect.ValueOf(iterator).IsNil() {
		if isFirst {
			fmt.Fprintf(writer, "%v", iterator.data)
			isFirst = false
		} else {
			fmt.Fprintf(writer, " %v", iterator.data)
		}
		iterator = iterator.next
	}
	fmt.Fprintln(writer)
}
