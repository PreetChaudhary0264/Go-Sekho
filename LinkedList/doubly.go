package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (ll *LinkedList[T]) insert(val T) {
	newNode := Node[T]{data: val}

	if ll.head == nil {
		ll.head = &newNode
		ll.tail = &newNode
		return
	}

	temp := ll.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &newNode
	newNode.prev = temp
	ll.tail = &newNode
}

func (ll *LinkedList[T]) print() {
	temp := ll.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Print(nil)
}

func (ll *LinkedList[T]) printFromLast(){
	temp := ll.tail
	for temp != nil {
		fmt.Print(temp.data, " <- ")
		temp = temp.prev
	}
	fmt.Println(nil)
}

func main() {
    ll := LinkedList[int]{}
	ll.insert(1)
	ll.insert(2)
	ll.insert(3)
	ll.print()
	ll.printFromLast()
}