package mutable

import (
	"github.com/karlsing/data-structure/core/iterator"
	"github.com/karlsing/data-structure/core/list"
)

type ArrayList[T any] struct {
	arr []T
}

func (list *ArrayList[T]) Length() int {
	return len(list.arr)
}

func (list *ArrayList[T]) Get(index int) T {
	return list.arr[index]
}

func (list *ArrayList[T]) Set(index int, elem T) {
	list.arr[index] = elem
}

func (list *ArrayList[T]) Push(elem T) {
	list.arr = append(list.arr, elem)
}

func (list *ArrayList[T]) Concat(l list.List[T]) list.List[T] {
	temp := list.Copy()
	iterator.Foreach(l.Iter(), func(elem T) {
		temp.Push(elem)
	})
	return temp
}

func (list *ArrayList[T]) Copy() list.List[T] {
	temp := make([]T, 0, list.Length())
	copy(temp, list.arr)
	return &ArrayList[T]{
		arr: temp,
	}
}

func (list *ArrayList[T]) Sub(start, length, step int) list.List[T] {
	return &ArrayList[T]{
		arr: list.arr[start:length:step],
	}
}

func (list *ArrayList[T]) ToSlice() []T {
	return list.arr
}

func (list *ArrayList[T]) Iter() iterator.Iterable[T] {
	return &arrayListIter[T]{0, list}
}

type arrayListIter[T any] struct {
	pos  int
	list *ArrayList[T]
}

func (it *arrayListIter[T]) HasNext() bool {
	return it.list.Length() == it.pos
}

func (it *arrayListIter[T]) Next() T {
	e := it.list.Get(it.pos)
	it.pos++
	return e
}
