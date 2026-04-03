package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (ll *LinkedList[T]) insert(val T) {
	newNode := Node[T]{data: val}

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	temp := ll.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &newNode
}

func (ll *LinkedList[T]) print() {
	temp := ll.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Print(nil)
}

func main4() {
    ll := LinkedList[int]{}
	ll.insert(1)
	ll.insert(2)
	ll.insert(3)
	ll.print()
}