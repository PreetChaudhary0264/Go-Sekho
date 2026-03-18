package main

import "fmt"

type Iterator[T any] struct {
	data  []T
	index int
}

func newIterator[T any](data []T) *Iterator[T] {
	return &Iterator[T]{data: data, index: 0}
}

func (it *Iterator[T]) hasNext() bool {
	return it.index < len(it.data)
}

func (it *Iterator[T]) next() T {
	val := it.data[it.index]
	it.index++
	return val
}

func main() {
	it := newIterator([]int{1, 2, 3, 4, 5})
	for it.hasNext() {
		fmt.Println(it.next())
	}
}