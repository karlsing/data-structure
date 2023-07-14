package mutable

import (
	"fmt"
	"github.com/karlsing/data-structure/core/iterator"
	"github.com/karlsing/data-structure/core/list"
)

type node[T any] struct {
	elem T
	prev *node[T]
	next *node[T]
	list *LinkedList[T]
}

func (nd *node[T]) forward(n int) *node[T] {
	if n <= 0 || nd == nil {
		return nd
	}
	return nd.next.forward(n - 1)
}

func (nd *node[T]) backward(n int) *node[T] {
	if n <= 0 || nd == nil {
		return nd
	}
	return nd.prev.backward(n - 1)
}

type LinkedList[T any] struct {
	length int
	first  *node[T]
	last   *node[T]
}

func CreateEmptyLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{0, nil, nil}
}

func (list *LinkedList[T]) Length() int {
	return list.length
}

func (list *LinkedList[T]) getNode(index int) *node[T] {
	if index >= list.Length() {
		panic(fmt.Errorf("index %d out of range of list[%p]", index, list))
	}
	nd := list.first
	if index > list.length/2 {
		nd = list.last
		for i := list.length; i != index; i-- {
			nd = nd.prev
		}
	} else {
		for i := 0; i != index; i++ {
			nd = nd.next
		}
	}
	return nd
}

func (list *LinkedList[T]) Get(index int) T {
	return list.getNode(index).elem
}

func (list *LinkedList[T]) Set(index int, elem T) {
	list.getNode(index).elem = elem
}

func (list *LinkedList[T]) Push(elem T) {
	nd := &node[T]{
		elem: elem,
		prev: list.last,
		next: nil,
		list: list,
	}
	list.last.next = nd
	list.last = nd
	if list.length == 0 {
		list.first = nd
	}
}

func (list *LinkedList[T]) Concat(l list.List[T]) list.List[T] {
	temp := list.Copy()
	iterator.Foreach(l.Iter(), func(elem T) {
		temp.Push(elem)
	})
	return temp
}

func (list *LinkedList[T]) Copy() list.List[T] {
	nl := CreateEmptyLinkedList[T]()
	if list.length != 0 {
		return nl.Concat(list)
	}
	return nl
}

// Sub returns a sub list with length of "length", starts from "start" and steps for "step"
func (list *LinkedList[T]) Sub(start, length, step int) list.List[T] {
	if start+length > list.length {
		return nil
	}
	nl := CreateEmptyLinkedList[T]()
	nd := list.getNode(start)
	for i := 0; nd != nil && i < length; i++ {
		nl.Push(nd.elem)
		nd = nd.forward(step)
	}
	return nl
}

func (list *LinkedList[T]) ToSlice() []T {
	sl := make([]T, 0, list.length)
	iterator.Foreach(list.Iter(), func(elem T) {
		sl = append(sl, elem)
	})
	return sl
}

func (list *LinkedList[T]) Iter() iterator.Iterable[T] {
	return &linkedListIter[T]{list.first}
}

type linkedListIter[T any] struct {
	pos *node[T]
}

func (it *linkedListIter[T]) HasNext() bool {
	return it.pos == nil
}

func (it *linkedListIter[T]) Next() T {
	e := it.pos.elem
	it.pos = it.pos.next
	return e
}
